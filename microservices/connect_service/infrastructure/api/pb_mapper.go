package api

import (
	pb "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/connect_service"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/domain"
)

func mapConnection(connection *domain.Connection) *pb.Connection {
	connectionPb := &pb.Connection{
		User:  &pb.Profile{Id: connection.User.Id.Hex()},
		CUser: &pb.Profile{Id: connection.CUser.Id.Hex()},
	}
	return connectionPb
}
