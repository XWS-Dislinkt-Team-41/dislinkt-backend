package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/api_gateway/domain"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/api_gateway/infrastructure/services"
	connections "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/connect_service"
	userService "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type ConnectHandler struct {
	connectClientAddress string
	userClientAddress    string
}

func NewConnectHandler(connectClientAddress, userClientAddress string) Handler {
	return &ConnectHandler{
		connectClientAddress: connectClientAddress,
		userClientAddress:    userClientAddress,
	}
}

func (handler *ConnectHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/user/{userId}/connect", handler.Connect)
	if err != nil {
		panic(err)
	}
}

func (handler *ConnectHandler) Connect(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	userId := pathParams["userId"]
	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var cUserId string
	json.NewDecoder(r.Body).Decode(&cUserId)

	isUserPrivate, err := handler.IsUserPrivate(cUserId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var connection *domain.Connection
	if *isUserPrivate {
		connection, err = handler.ConnectRequest(userId, cUserId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		connection, err = handler.InviteRequest(userId, cUserId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response, err := json.Marshal(connection)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (handler *ConnectHandler) ConnectRequest(userId, cUserId string) (*domain.Connection, error) {
	connectClient := services.NewConnectClient(handler.connectClientAddress)
	r, err := connectClient.Connect(context.TODO(), &connections.ConnectRequest{UserId: userId, CUser: &connections.Profile{Id: cUserId}})
	if err != nil {
		return nil, err
	}
	connection := &domain.Connection{
		User:  domain.Profile{Id: r.Connection.User.Id},
		CUser: domain.Profile{Id: r.Connection.CUser.Id},
	}
	return connection, nil
}

func (handler *ConnectHandler) InviteRequest(userId, cUserId string) (*domain.Connection, error) {
	connectClient := services.NewConnectClient(handler.connectClientAddress)
	r, err := connectClient.Invite(context.TODO(), &connections.InviteRequest{UserId: userId, CUser: &connections.Profile{Id: cUserId}})
	if err != nil {
		return nil, err
	}
	connection := &domain.Connection{
		User:  domain.Profile{Id: r.Connection.User.Id},
		CUser: domain.Profile{Id: r.Connection.CUser.Id},
	}
	return connection, nil
}

func (handler *ConnectHandler) IsUserPrivate(userId string) (*bool, error) {
	userClient := services.NewUserClient(handler.userClientAddress)
	userProfileStatus, err := userClient.IsPrivate(context.TODO(), &userService.IsPrivateRequest{Id: userId})
	if err != nil {
		return nil, err
	}
	return &userProfileStatus.Private, nil
}
