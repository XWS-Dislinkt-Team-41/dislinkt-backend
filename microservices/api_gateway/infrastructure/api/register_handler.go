package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/api_gateway/domain"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/api_gateway/infrastructure/services"
	auth "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/auth_service"
	conn "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/connect_service"
	user "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegisterHandler struct {
	authClientAddress    string
	userClientAddress    string
	connectClientAddress string
}

func NewRegisterHandler(userClientAddress, authClientAddress, connectClientAddress string) Handler {
	return &RegisterHandler{
		userClientAddress:    userClientAddress,
		authClientAddress:    authClientAddress,
		connectClientAddress: connectClientAddress,
	}
}

func (handler *RegisterHandler) Init(mux *runtime.ServeMux) {

	err := mux.HandlePath("POST", "/user/register", handler.Register)
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
			Id:           userRequest.Id,
			Username:     userRequest.Username,
			Password:     userRequest.Password,
			IsPrivate:    userRequest.IsPrivate,
			Firstname:    userRequest.Firstname,
			Lastname:     userRequest.Lastname,
			MobileNumber: userRequest.MobileNumber,
			Email:        userRequest.Email,
		},
		UserCredential: auth.UserCredential{
			Username: userRequest.Username,
			Password: userRequest.Password,
		},
		Profile: conn.Profile{
			Id:      userRequest.Id,
			Private: userRequest.IsPrivate,
		},
	}

	err = handler.RegisterUser(registerRequest)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	} else {
		err = handler.RegisterUserCredential(registerRequest)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		err = handler.RegisterProfile(registerRequest)
		if err != nil {
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
}

func (handler *RegisterHandler) RegisterUser(registerUserRequest *domain.RegisterRequest) error {
	userClient := services.NewUserClient(handler.userClientAddress)
	user, err := userClient.Register(context.TODO(), &user.RegisterRequest{User: &registerUserRequest.User})
	registerUserRequest.User = *user.User
	if err != nil {
		return err
	}
	if user == nil {
		return status.Error(codes.AlreadyExists, "User already exists with same credentials")
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

func (handler *RegisterHandler) RegisterProfile(registerUserRequest *domain.RegisterRequest) error {
	connectClient := services.NewConnectClient(handler.connectClientAddress)
	registerUserRequest.Profile.Id = registerUserRequest.User.Id
	registerUserRequest.Profile.Private = registerUserRequest.User.IsPrivate
	_, err := connectClient.Register(context.TODO(), &conn.RegisterRequest{User: &registerUserRequest.Profile})
	if err != nil {
		return err
	}
	return nil
}
