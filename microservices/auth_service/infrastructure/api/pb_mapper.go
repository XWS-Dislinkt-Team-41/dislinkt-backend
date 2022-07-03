package api

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/auth_service/domain"
	pb "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/auth_service"
	events "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/saga/register_user"
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

func mapUserDetails(user *events.UserDetails) *pb.UserDetails {
	userPb := &pb.UserDetails{
		Id:           user.Id,
		Username:     user.Username,
		Password:     user.Password,
		IsPrivate:    user.IsPrivate,
		Firstname:    user.Firstname,
		Lastname:     user.Lastname,
		Email:        user.Email,
		MobileNumber: user.MobileNumber,
	}
	return userPb
}

func mapPbUserDetails(userPb *pb.UserDetails) *events.UserDetails {
	user := &events.UserDetails{
		Username:     userPb.Username,
		Password:     userPb.Password,
		IsPrivate:    userPb.IsPrivate,
		Firstname:    userPb.Firstname,
		Lastname:     userPb.Lastname,
		Email:        userPb.Email,
		MobileNumber: userPb.MobileNumber,
	}
	return user
}
