package api

import (
	saga "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/saga/messaging"
	events "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/saga/make_account_private"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service/application"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChangePrivacyCommandHandler struct {
	userService       *application.UserService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewChangePrivacyCommandHandler(userService *application.UserService, publisher saga.Publisher, subscriber saga.Subscriber) (*ChangePrivacyCommandHandler, error) {
	o := &ChangePrivacyCommandHandler{
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

func (handler *ChangePrivacyCommandHandler) handle(command *events.ChangePrivacyCommand) {
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

	reply := events.ChangePrivacyReply{User: command.User}

	switch command.Type {
	case events.ChangePrivacy:
		_, err := handler.userService.MakeAccountPrivate(user.id)
		if err != nil {
			reply.Type = events.AccountNotPrivated
			break
		}
		reply.Type = events.AccountPrivated
	case events.RollbackChangePrivacy:
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
