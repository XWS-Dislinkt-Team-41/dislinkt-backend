package api

import (
	"context"
	"fmt"
	"os"

	pb "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/user_service"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service/application"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	user, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	userPb := mapUser(user)
	response := &pb.GetResponse{
		User: userPb,
	}
	return response, nil
}

func (handler *UserHandler) GetPrincipal(ctx context.Context, request *pb.SearchPublicRequest) (*pb.GetResponse, error) {
	fmt.Println(request.Filter)
	user, err := handler.service.GetPrincipal(request.Filter)

	if err != nil {
		return nil, err
	}
	userPb := mapUser(user)
	response := &pb.GetResponse{
		User: userPb,
	}
	return response, nil
}

func (handler *UserHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	users, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Users: []*pb.User{},
	}
	for _, user := range users {
		current := mapUser(user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (handler *UserHandler) GetAllPublicUserId(ctx context.Context, request *pb.GetAllPublicUserIdRequest) (*pb.GetAllPublicUserIdResponse, error) {
	ids, err := handler.service.GetAllPublicUserId()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllPublicUserIdResponse{
		Ids: []string{},
	}
	for _, id := range ids {
		current := id.Hex()
		response.Ids = append(response.Ids, current)
	}
	return response, nil
}

func (handler *UserHandler) IsPrivate(ctx context.Context, request *pb.IsPrivateRequest) (*pb.IsPrivateResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	isPrivate, err := handler.service.IsPrivate(objectId)
	if err != nil {
		return nil, err
	}
	response := &pb.IsPrivateResponse{
		Private: isPrivate,
	}
	return response, nil
}

func (handler *UserHandler) SearchPublic(ctx context.Context, request *pb.SearchPublicRequest) (*pb.SearchPublicResponse, error) {
	filter := request.Filter
	users, err := handler.service.SearchPublic(filter)
	if err != nil {
		return nil, err
	}
	response := &pb.SearchPublicResponse{
		Users: []*pb.User{},
	}
	for _, user := range users {
		current := mapUser(user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (handler *UserHandler) Search(ctx context.Context, request *pb.SearchPublicRequest) (*pb.SearchPublicResponse, error) {
	filter := request.Filter
	users, err := handler.service.Search(filter)
	if err != nil {
		return nil, err
	}
	response := &pb.SearchPublicResponse{
		Users: []*pb.User{},
	}
	for _, user := range users {
		current := mapUser(user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (handler *UserHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	userRequest := mapNewUser(request.User)
	user, err := handler.service.Register(userRequest)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, status.Error(codes.AlreadyExists, "User already exists with same credentials")
	}
	response := &pb.RegisterResponse{
		User: mapRegisterUser(user),
	}
	return response, err
}

func (handler *UserHandler) UpdatePersonalInfo(ctx context.Context, request *pb.UpdatePersonalInfoRequest) (*pb.UpdatePersonalInfoResponse, error) {
	user := mapPersonalInfoUser(request.User)
	user, err := handler.service.UpdatePersonalInfo(user)
	if err != nil {
		return nil, err
	}
	userPb := mapUser(user)
	response := &pb.UpdatePersonalInfoResponse{
		User: userPb,
	}
	return response, err
}

func (handler *UserHandler) UpdateCareerInfo(ctx context.Context, request *pb.UpdateCareerInfoRequest) (*pb.UpdateCareerInfoResponse, error) {
	user := mapCareerInfoUser(request.User)
	user, err := handler.service.UpdateCareerInfo(user)
	if err != nil {
		return nil, err
	}
	userPb := mapUser(user)
	response := &pb.UpdateCareerInfoResponse{
		User: userPb,
	}
	return response, err
}

func (handler *UserHandler) UpdateInterestsInfo(ctx context.Context, request *pb.UpdateInterestsInfoRequest) (*pb.UpdateInterestsInfoResponse, error) {
	user := mapInterestsInfoUser(request.User)
	user, err := handler.service.UpdateInterestsInfo(user)
	if err != nil {
		return nil, err
	}
	userPb := mapUser(user)
	response := &pb.UpdateInterestsInfoResponse{
		User: userPb,
	}
	return response, err
}

func (handler *UserHandler) ChangeAccountPrivacy(ctx context.Context, request *pb.ChangeAccountPrivacyRequest) (*pb.ChangeAccountPrivacyResponse, error) {
	user := mapPbUserDetails(request.User)
	user, err := handler.service.ChangeAccountPrivacy(user)
	if err != nil {
		return nil, err
	}
	userPb := mapUserDetails(user)
	response := &pb.ChangeAccountPrivacyResponse{
		User: userPb,
	}
	return response, err
}

func getUsernameFromJWT(ctx context.Context) string {
	username := ""
	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Println("Uso " + md["Authorization"][0])
	jwt.Parse(md["Authorization"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(("Invalid Signing Method"))
		}
		aud := "billing.jwtgo.io"
		checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
		if !checkAudience {
			return nil, fmt.Errorf(("invalid aud"))
		}
		// verify iss claim
		iss := "jwtgo.io"
		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return nil, fmt.Errorf(("invalid iss"))
		}
		fmt.Println("Usernameeeeee")
		username = token.Claims.(jwt.MapClaims)["username"].(string)
		fmt.Println("Usernamee" + username)
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	return username
}
