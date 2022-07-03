package api

import (
	saga "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/saga/messaging"
	events "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/saga/register_user"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/application"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterUserCommandHandler struct {
	connectService    *application.ConnectService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewRegisterUserCommandHandler(connectService *application.ConnectService, publisher saga.Publisher, subscriber saga.Subscriber) (*RegisterUserCommandHandler, error) {
	o := &RegisterUserCommandHandler{
		connectService:    connectService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *RegisterUserCommandHandler) handle(command *events.RegisterUserCommand) {
	id, err := primitive.ObjectIDFromHex(command.User.Id)
	if err != nil {
		return
	}
	user := &domain.Profile{
		Id:      id,
		Private: command.User.IsPrivate,
	}

	reply := events.RegisterUserReply{User: command.User}

	switch command.Type {
	case events.RegisterUserNode:
		_, err = handler.connectService.Register(*user)
		if err != nil {
			reply.Type = events.UserNodeNotRegistered
			break
		}
		reply.Type = events.UserNodeRegistered
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
