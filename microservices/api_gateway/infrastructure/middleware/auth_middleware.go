package middleware

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"context"
	
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/api_gateway/infrastructure/services"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/api_gateway/infrastructure/api"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func IsAuthenticated(handler *runtime.ServeMux) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isProtectedRoute(r.Method, r.URL.Path) {
			if r.Header["Authorization"] != nil {

				authEndpoint := fmt.Sprintf("%s:%s", "auth_service", "8000")
				authHandler := api.NewAuthandler(authEndpoint)
				authHandler.Init(handler)
				authClient := services.NewAuthClient(handler.authClientAddress)
				//auth client
				response, err := authClient.RBAC(context.TODO(),"leka",r.URL.Path,r.Method)
				if err != nil {
					fmt.Fprintf(w, err.Error())
					return 
				}
				if !response.Response{
					fmt.Fprintf(w, "Endpoint access denied")
					return
				}
				//valid_route, err := auth.RBAC(claims.username,r.URL.Path,r.method)
				// if  err return nil,err
				// //joboffer gadja iz agentske
				// if r.urlPath {  token.Claims["apiToken"] }
				// /connections POST leka
				// mongodb role_permissiondb
				// Collection roles
				// username role
				// leka     user
				// pape     owner
				// dare     admin
				//Collection permission
				// role endpoint method
				// admin /users  GET
				// admin /users  POST

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
		if path == "/user/register" ||
			path == "/auth/login" {
			return false
		}
	}
	return true
}
