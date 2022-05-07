package api

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/auth_service/domain"
	pb "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/auth_service"
)

func mapUserCredential(userCredential *domain.UserCredential) *pb.UserCredential {
	userCredentialPb := &pb.UserCredential{
		Username: userCredential.Username,
		Password: userCredential.Password,
	}
	return userCredentialPb
}

func mapPbUserCredential(userCredential *pb.UserCredential) *domain.UserCredential {
	userCredentialPb := &domain.UserCredential{
		Username: userCredential.Username,
		Password: userCredential.Password,
	}
	return userCredentialPb
}
