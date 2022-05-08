package api

import (
	pb "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/connect_service"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapConnection(connection *domain.Connection) *pb.Connection {
	if connection == nil {
		return nil
	}
	connectionPb := &pb.Connection{
		User:  &pb.Profile{Id: connection.User.Id.Hex()},
		CUser: &pb.Profile{Id: connection.CUser.Id.Hex()},
	}
	return connectionPb
}

func mapProfilePb(profile *pb.Profile) (*domain.Profile, error) {
	userId, err := primitive.ObjectIDFromHex(profile.Id)
	if err != nil {
		return nil, err
	}
	user := &domain.Profile{
		Id:      userId,
		Private: profile.Private,
	}
	return user, nil
}

func mapProfile(profile *domain.Profile) *pb.Profile {
	if profile == nil {
		return nil
	}
	profilePb := &pb.Profile{
		Id:      profile.Id.Hex(),
		Private: profile.Private,
	}
	return profilePb
}
