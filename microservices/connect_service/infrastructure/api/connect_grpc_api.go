package api

import (
	"context"

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

func (handler *ConnectHandler) Connect(ctx context.Context, request *pb.ConnectRequest) (*pb.EmptyResponse, error) {
	err := handler.service.Connect(request.User, request.UserConnect)
	if err != nil {
		return nil, err
	}
	response := &pb.EmptyResponse{}
	return response, nil
}

func (handler *ConnectHandler) UnConnect(ctx context.Context, request *pb.UnConnectRequest) (*pb.EmptyResponse, error) {
	err := handler.service.UnConnect(request.User, request.UserConnect)
	if err != nil {
		return nil, err
	}
	response := &pb.EmptyResponse{}
	return response, nil
}

func (handler *ConnectHandler) GetUserConnections(ctx context.Context, request *pb.GetUserConnectionsRequest) (*pb.GetUserConnectionsResponse, error) {
	connections, err := handler.service.GetUserConnections(request.Id)
	if err != nil {
		return nil, err
	}
	response := &pb.GetUserConnectionsResponse{
		Connections: connections,
	}
	return response, nil
}
