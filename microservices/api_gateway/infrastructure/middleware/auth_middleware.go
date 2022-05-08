package middleware

import (
	"fmt"
	"net/http"
	"os"
	"regexp"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func IsAuthenticated(handler *runtime.ServeMux) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isProtectedRoute(r.Method, r.URL.Path) {
			if r.Header["Authorization"] != nil {

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