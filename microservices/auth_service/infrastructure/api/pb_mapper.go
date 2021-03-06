package api

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/auth_service/domain"
	pb "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/auth_service"
	events "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/saga/register_user"
)

func mapUserRole(status domain.Role) pb.UserCredential_Role {
	switch status {
	case domain.USER:
		return pb.UserCredential_USER
	case domain.AGENT:
		return pb.UserCredential_AGENT
	}
	return pb.UserCredential_ADMIN
}

func mapNewUserRole(status pb.UserCredential_Role) domain.Role {
	switch status {
	case pb.UserCredential_USER:
		return domain.USER
	case pb.UserCredential_AGENT:
		return domain.AGENT
	}
	return domain.ADMIN
}

func mapUserDetailsRole(status events.Role) pb.UserDetails_Role {
	switch status {
	case events.USER:
		return pb.UserDetails_USER
	}
	return pb.UserDetails_ADMIN
}

func mapNewUserDetailsRole(status pb.UserDetails_Role) events.Role {
	switch status {
	case pb.UserDetails_USER:
		return events.USER
	}
	return events.ADMIN
}

func mapPermissionRole(status domain.Role) pb.Permission_Role {
	switch status {
	case domain.USER:
		return pb.Permission_USER
	case domain.AGENT:
		return pb.Permission_AGENT
	}
	return pb.Permission_ADMIN
}

func mapNewPermissionRole(status pb.Permission_Role) domain.Role {
	switch status {
	case pb.Permission_USER:
		return domain.USER
	case pb.Permission_AGENT:
		return domain.AGENT
	}
	return domain.ADMIN
}

func mapMethod(status domain.Method) pb.Permission_Method {
	switch status {
	case domain.POST:
		return pb.Permission_POST
	case domain.GET:
		return pb.Permission_GET
	case domain.PUT:
		return pb.Permission_PUT
	}
	return pb.Permission_DELETE
}

func mapNewMethod(status pb.Permission_Method) domain.Method {
	switch status {
	case pb.Permission_POST:
		return domain.POST
	case pb.Permission_GET:
		return domain.GET
	case pb.Permission_PUT:
		return domain.PUT
	}
	return domain.DELETE
}

func mapUserCredential(userCredential *domain.UserCredential) *pb.UserCredential {
	userCredentialPb := &pb.UserCredential{
		Username: userCredential.Username,
		Password: userCredential.Password,
		Role:     mapUserRole(userCredential.Role),
	}
	return userCredentialPb
}

func mapPermission(permission *domain.Permission) *pb.Permission {
	permissionPb := &pb.Permission{
		Role:   mapPermissionRole(permission.Role),
		Method: mapMethod(permission.Method),
		Url:    permission.Url,
	}
	return permissionPb
}

func mapPbUserCredential(userCredential *pb.UserCredential) *domain.UserCredential {
	userCredentialPb := &domain.UserCredential{
		Username: userCredential.Username,
		Password: userCredential.Password,
		Role:     mapNewUserRole(userCredential.Role),
	}
	return userCredentialPb
}

func mapPbPermission(permissionPb *pb.Permission) *domain.Permission {
	permission := &domain.Permission{
		Role:   mapNewPermissionRole(permissionPb.Role),
		Method: mapNewMethod(permissionPb.Method),
		Url:    permissionPb.Url,
	}
	return permission
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
		Role:         mapUserDetailsRole(user.Role),
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
		Role:         mapNewUserDetailsRole(userPb.Role),
	}
	return user
}
