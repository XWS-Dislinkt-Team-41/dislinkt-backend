package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/api_gateway/domain"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/api_gateway/infrastructure/services"
	user "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/user_service"
	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type GetCurrentUserHandler struct {
	userClientAddress string
}

func NewGetCurrentUserHandler(userClientAddress string) Handler {
	return &GetCurrentUserHandler{
		userClientAddress: userClientAddress,
	}
}

func (handler *GetCurrentUserHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/user/principal", handler.GetCurrentUser)
	if err != nil {
		panic(err)
	}
}

func (handler *GetCurrentUserHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	token := r.Header["Authorization"][0]
	claims, _ := extractClaims(token)
	username := claims["username"].(string)
	principalRequest := &domain.PrincipalRequest{Username: username}

	err := handler.GetPrincipalUser(principalRequest)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	response, err := json.Marshal(principalRequest.User)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (handler *GetCurrentUserHandler) GetPrincipalUser(principal *domain.PrincipalRequest) error {
	userClient := services.NewUserClient(handler.userClientAddress)
	res, err := userClient.GetPrincipal(context.TODO(), &user.SearchPublicRequest{Filter: principal.Username})
	principal.User = *res.User
	if err != nil {
		return err
	}
	return nil
}

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
