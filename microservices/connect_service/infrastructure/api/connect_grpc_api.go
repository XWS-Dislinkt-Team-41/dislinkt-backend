package api

import (
	pb "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/connect_service"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/application"
)

type ConnectHandler struct {
	pb.UnimplementedConnectServiceServer
	service *application.ConnectService
}

func NewConnectHandler(service *application.ConnectService) *ConnectHandler {
	return &ConnectHandler{
		service: service,
	}
}
