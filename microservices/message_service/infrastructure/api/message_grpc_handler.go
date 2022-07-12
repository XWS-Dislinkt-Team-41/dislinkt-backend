package api

import (
	"context"

	pb "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/message_service"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/message_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageHandler struct {
	pb.UnimplementedMessageServiceServer
	service *application.MessageService
}

func NewMessageHandler(service *application.MessageService) *MessageHandler {
	return &MessageHandler{
		service: service,
	}
}

func (handler *MessageHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	connectedId, err := primitive.ObjectIDFromHex(request.ConnectedId)
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	messages, err := handler.service.Get(objectId, connectedId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetResponse{
		Messages: []*pb.Message{},
	}

	for _, Message := range messages {
		current := mapMessage(Message)
		response.Messages = append(response.Messages, current)
	}
	return response, nil
}
