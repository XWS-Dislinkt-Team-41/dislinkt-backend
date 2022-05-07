package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/api_gateway/domain"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/api_gateway/infrastructure/services"
	auth "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/auth_service"
	user "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type RegisterHandler struct {
	authClientAddress string
	userClientAddress string
}

func NewRegisterHandler(userClientAddress, authClientAddress string) Handler {
	return &RegisterHandler{
		userClientAddress: userClientAddress,
		authClientAddress: authClientAddress,
	}
}

func (handler *RegisterHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/user/comporeg", handler.Register)
	if err != nil {
		panic(err)
	}
}

func (handler *RegisterHandler) Register(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

	// ovde iz request-a
	var userRequest user.User

	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	registerRequest := &domain.RegisterRequest{
		User: user.User{
			Id:       userRequest.Id,
			Username: userRequest.Username,
			Password: userRequest.Password,
		},
		UserCredential: auth.UserCredential{
			Username: userRequest.Username,
			Password: userRequest.Password,
		},
	}

	err = handler.RegisterUser(registerRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		err1 := handler.RegisterUserCredential(registerRequest)
		if err1 != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		response, err := json.Marshal(registerRequest.User)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
	return
}

func (handler *RegisterHandler) RegisterUser(registerUserRequest *domain.RegisterRequest) error {
	userClient := services.NewUserClient(handler.userClientAddress)
	_, err := userClient.Register(context.TODO(), &user.RegisterRequest{User: &registerUserRequest.User})
	if err != nil {
		return err
	}
	return nil
}

func (handler *RegisterHandler) RegisterUserCredential(registerUserRequest *domain.RegisterRequest) error {
	authClient := services.NewAuthClient(handler.authClientAddress)
	_, err := authClient.Register(context.TODO(), &auth.RegisterRequest{User: &registerUserRequest.UserCredential})
	if err != nil {
		return err
	}
	return nil
}
