package api

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/auth_service/domain"
	pb "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/auth_service"
)

func mapUserCredential(userCredential *domain.UserCredential) *pb.UserCredential {
	userCredentialPb := &pb.UserCredential{
		Username: userCredential.Username,
		Password: userCredential.Password,
		Role: userCredential.Role,
	}
	return userCredentialPb
}

func mapPermission(permission *domain.Permission) *pb.Permission {
	permissionPb := &pb.Permission{
		Role: permission.Username,
		Method: permission.Method,
		Url: permission.Url,
	}
	return permissionPb
}

func mapPbUserCredential(userCredential *pb.UserCredential) *domain.UserCredential {
	userCredentialPb := &domain.UserCredential{
		Username: userCredential.Username,
		Password: userCredential.Password,
		Role: userCredential.Role,
	}
	return userCredentialPb
}

func mapPbPermission(permissionPb *pb.Permission) *domain.Permission {
	permission := &domain.Permission{
		Role: permissionPb.Username,
		Method: permissionPb.Method,
		Url: permissionPb.Url,
	}
	return permission
}