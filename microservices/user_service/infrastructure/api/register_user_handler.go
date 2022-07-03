package api

import (
	saga "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/saga/messaging"
	events "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/saga/register_user"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service/application"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterUserCommandHandler struct {
	userService       *application.UserService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewRegisterUserCommandHandler(userService *application.UserService, publisher saga.Publisher, subscriber saga.Subscriber) (*RegisterUserCommandHandler, error) {
	o := &RegisterUserCommandHandler{
		userService:       userService,
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
	user := &domain.User{
		Id:           id,
		Firstname:    command.User.Firstname,
		Lastname:     command.User.Lastname,
		Email:        command.User.Email,
		MobileNumber: command.User.MobileNumber,
		Username:     command.User.Username,
		Password:     command.User.Password,
		IsPrivate:    command.User.IsPrivate,
	}

	reply := events.RegisterUserReply{User: command.User}

	switch command.Type {
	case events.RegisterUser:
		_, err := handler.userService.Register(user)
		if err != nil {
			reply.Type = events.UserNotRegistered
			break
		}
		reply.Type = events.UserRegistered
	case events.RollbackRegisterUser:
		err := handler.userService.DeleteById(user.Id)
		if err != nil {
			return
		}
		reply.Type = events.UserRolledBack
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
