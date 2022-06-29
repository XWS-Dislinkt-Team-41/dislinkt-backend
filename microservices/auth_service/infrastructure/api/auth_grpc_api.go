package api

import (
	"context"

	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/auth_service/application"
	pb "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/auth_service"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	service *application.AuthService
}

func NewAuthHandler(service *application.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (handler *AuthHandler) RBAC(ctx context.Context, request *pb.RBACRequest) (*pb.RBACResponse, error) {
	userCredential := mapPbUserCredential(request.User)
	permissionCredential := mapPbPermission(request.Permission)
	response, err := handler.service.RBAC(userCredential.Username,permissionCredential.Method,permissionCredential.Url)
	if err != nil {
		return nil, err
	}
	return &pb.RBACResponse{
		Response: response,
	}, nil
}

func (handler *AuthHandler) Login(ctx context.Context, request *pb.LoginRequest) (*pb.JWTResponse, error) {
	userCredential := mapPbUserCredential(request.User)
	jwt, err := handler.service.Login(userCredential)
	if err != nil {
		return nil, err
	}
	return &pb.JWTResponse{
		Token: jwt.Token,
	}, nil
}

func (handler *AuthHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	userCredential := mapPbUserCredential(request.User)
	user, err := handler.service.Register(userCredential)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{
		User: mapUserCredential(user),
	}, nil
}
