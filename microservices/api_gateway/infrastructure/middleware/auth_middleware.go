package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/api_gateway/infrastructure/services"
	auth "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/auth_service"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func ConvertStringToMethod(method string) auth.Permission_Method {
	if method == "POST" {
		return auth.Permission_POST
	} else if method == "GET" {
		return auth.Permission_GET
	} else if method == "PUT" {
		return auth.Permission_PUT
	} else {
		return auth.Permission_DELETE
	}
}

func IsAuthenticated(handler *runtime.ServeMux) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isProtectedRoute(r.Method, r.URL.Path) {
			if r.Header["Authorization"] != nil {
				authEndpoint := fmt.Sprintf("%s:%s", "auth_service", "8000")
				authClient := services.NewAuthClient(authEndpoint)

				token, err := jwt.Parse(r.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf(("Invalid Signing Method"))
					}
					aud := "billing.jwtgo.io"
					checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
					if !checkAudience {
						return nil, fmt.Errorf(("invalid aud"))
					}
					// verify iss claim
					iss := "jwtgo.io"
					checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
					if !checkIss {
						return nil, fmt.Errorf(("invalid iss"))
					}
					username := token.Claims.(jwt.MapClaims)["username"]
					usernameStr := fmt.Sprintf("%v", username)
					tokenType := token.Claims.(jwt.MapClaims)["type"]

					if tokenType == "API" {
						response, err := authClient.RBAC(context.TODO(), &auth.RBACRequest{User: &auth.UserCredential{Username: usernameStr}, Permission: &auth.Permission{Method: ConvertStringToMethod(r.Method), Url: r.URL.Path}})
						if err != nil {
							return nil, err
						}
						if !response.Response {
							return nil, fmt.Errorf(("Endpoint access denied"))
						}
					} else {
						response, err := authClient.RBAC(context.TODO(), &auth.RBACRequest{User: &auth.UserCredential{Username: usernameStr}, Permission: &auth.Permission{Method: ConvertStringToMethod(r.Method), Url: r.URL.Path}})
						if err != nil {
							return nil, err
						}
						if !response.Response {
							return nil, fmt.Errorf(("Endpoint access denied"))
						}
					}

					return []byte(os.Getenv("JWT_SECRET_KEY")), nil
				})
				if err != nil {
					fmt.Fprintf(w, err.Error())
					return
				}

				if token.Valid {
					handler.ServeHTTP(w, r)
					return
				}

			} else {
				fmt.Fprintf(w, "No Authorization Token provided")
				return
			}
		}
		handler.ServeHTTP(w, r)
	})
}

func isProtectedRoute(method, path string) bool {
	publicProfilePostsPath, _ := regexp.MatchString(`\/user\/[0-9a-f]{24}\/public`, path)
	if method == "GET" {
		if path == "/post/public" ||
			publicProfilePostsPath {
			return false
		}
	}

	if method == "POST" {
		if path == "/auth/register" ||
			path == "/auth/login" ||
			path == "/jobOffer/search" ||
			path == "/user/search" ||
			path == "/auth/connectAgent" {
			return false
		}
	}
	return true
}
