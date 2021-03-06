package api

import (
	"context"

	pb "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/connect_service"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (handler *ConnectHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.ProfileResponse, error) {
	profile, err := mapProfilePb(request.User)
	if err != nil {
		return nil, err
	}
	user, err := handler.service.Register(*profile)
	if err != nil {
		return nil, err
	}
	response := &pb.ProfileResponse{
		User: mapProfile(user),
	}
	return response, nil
}

func (handler *ConnectHandler) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.ProfileResponse, error) {
	profile, err := mapProfilePb(request.User)
	if err != nil {
		return nil, err
	}
	user, err := handler.service.UpdateUser(*profile)
	if err != nil {
		return nil, err
	}
	response := &pb.ProfileResponse{
		User: mapProfile(user),
	}
	return response, nil
}

func (handler *ConnectHandler) Connect(ctx context.Context, request *pb.ConnectRequest) (*pb.ConnectionResponse, error) {
	userId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}
	cUserId, err := primitive.ObjectIDFromHex(request.CUser.Id)
	if err != nil {
		return nil, err
	}
	connection, err := handler.service.Connect(userId, cUserId)
	if err != nil {
		return nil, err
	}
	response := &pb.ConnectionResponse{
		Connection: mapConnection(connection),
	}
	return response, nil
}

func (handler *ConnectHandler) UnConnect(ctx context.Context, request *pb.UnConnectRequest) (*pb.EmptyRespones, error) {
	userId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}
	cUserId, err := primitive.ObjectIDFromHex(request.CUserId)
	if err != nil {
		return nil, err
	}
	err = handler.service.UnConnect(userId, cUserId)
	if err != nil {
		return nil, err
	}
	response := &pb.EmptyRespones{}
	return response, nil
}

func (handler *ConnectHandler) GetUserConnections(ctx context.Context, request *pb.GetUserConnectionsRequest) (*pb.GetUserConnectionsResponse, error) {
	userId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}
	connections, err := handler.service.GetUserConnections(userId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetUserConnectionsResponse{
		Connections: []*pb.Connection{},
	}
	for _, Connect := range connections {
		current := mapConnection(Connect)
		response.Connections = append(response.Connections, current)
	}
	return response, nil
}

func (handler *ConnectHandler) AcceptInvitation(ctx context.Context, request *pb.AcceptInvitationRequest) (*pb.ConnectionResponse, error) {
	userId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}
	cUserId, err := primitive.ObjectIDFromHex(request.CUserId)
	if err != nil {
		return nil, err
	}
	connection, err := handler.service.AcceptInvitation(userId, cUserId)
	if err != nil {
		return nil, err
	}
	response := &pb.ConnectionResponse{
		Connection: mapConnection(connection),
	}
	return response, nil
}

func (handler *ConnectHandler) DeclineInvitation(ctx context.Context, request *pb.DeclineInvitationRequest) (*pb.EmptyRespones, error) {
	userId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}
	cUserId, err := primitive.ObjectIDFromHex(request.CUserId)
	if err != nil {
		return nil, err
	}
	err = handler.service.DeclineInvitation(userId, cUserId)
	if err != nil {
		return nil, err
	}
	response := &pb.EmptyRespones{}
	return response, nil
}

func (handler *ConnectHandler) CancelInvitation(ctx context.Context, request *pb.CancelInvitationRequest) (*pb.EmptyRespones, error) {
	userId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}
	cUserId, err := primitive.ObjectIDFromHex(request.CUserId)
	if err != nil {
		return nil, err
	}
	err = handler.service.CancelInvitation(userId, cUserId)
	if err != nil {
		return nil, err
	}
	response := &pb.EmptyRespones{}
	return response, nil
}

func (handler *ConnectHandler) GetAllInvitations(ctx context.Context, request *pb.GetAllUserInvitationsRequest) (*pb.GetAllInvitationsResponse, error) {
	userId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}
	invites, err := handler.service.GetAllInvitations(userId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllInvitationsResponse{
		ConnectInvitations: []*pb.Connection{},
	}
	for _, Invite := range invites {
		current := mapConnection(Invite)
		response.ConnectInvitations = append(response.ConnectInvitations, current)
	}
	return response, nil
}

func (handler *ConnectHandler) GetAllSentInvitations(ctx context.Context, request *pb.GetAllSentInvitationsRequest) (*pb.GetAllInvitationsResponse, error) {
	userId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}
	invites, err := handler.service.GetAllSentInvitations(userId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllInvitationsResponse{
		ConnectInvitations: []*pb.Connection{},
	}
	for _, Invite := range invites {
		current := mapConnection(Invite)
		response.ConnectInvitations = append(response.ConnectInvitations, current)
	}
	return response, nil
}

func (handler *ConnectHandler) GetUserSuggestions(ctx context.Context, request *pb.GetUserSuggestionsRequest) (*pb.GetUserSuggestionsResponse, error) {
	userId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}
	userSuggestions, err := handler.service.GetUserSuggestions(userId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetUserSuggestionsResponse{
		UserSuggestions: []*pb.Profile{},
	}
	for _, user := range userSuggestions {
		current := mapProfile(user)
		response.UserSuggestions = append(response.UserSuggestions, current)
	}
	return response, nil
}

func (handler *ConnectHandler) Block(ctx context.Context, request *pb.BlockRequest) (*pb.BlockResponse, error) {
	userId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}
	bUserId, err := primitive.ObjectIDFromHex(request.BUser.Id)
	if err != nil {
		return nil, err
	}
	block, err := handler.service.Block(userId, bUserId)
	if err != nil {
		return nil, err
	}
	response := &pb.BlockResponse{
		Block: mapBlock(block),
	}
	return response, nil
}

func (handler *ConnectHandler) UnBlock(ctx context.Context, request *pb.UnBlockRequest) (*pb.EmptyRespones, error) {
	userId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}
	bUserId, err := primitive.ObjectIDFromHex(request.BUserId)
	if err != nil {
		return nil, err
	}
	err = handler.service.UnBlock(userId, bUserId)
	if err != nil {
		return nil, err
	}
	response := &pb.EmptyRespones{}
	return response, nil
}

func (handler *ConnectHandler) GetBlockedUsers(ctx context.Context, request *pb.GetBlockedUsersRequest) (*pb.GetBlockedUsersResponse, error) {
	userId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}
	blocks, err := handler.service.GetBlockedUsers(userId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetBlockedUsersResponse{
		Blocks: []*pb.Block{},
	}
	for _, Block := range blocks {
		current := mapBlock(Block)
		response.Blocks = append(response.Blocks, current)
	}
	return response, nil
}
