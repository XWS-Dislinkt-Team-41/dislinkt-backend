package api

import (
	"fmt"

	events "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/saga/change_account_privacy"
	saga "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/saga/messaging"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service/application"
	//"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service/domain"
	//"go.mongodb.org/mongo-driver/bson/primitive"
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
	// id, err := primitive.ObjectIDFromHex(command.User.Id)
	// if err != nil {
	// 	return
	// }
	// user := &domain.User{
	// 	Id:           id,
	// 	Firstname:    command.User.Firstname,
	// 	Lastname:     command.User.Lastname,
	// 	Email:        command.User.Email,
	// 	MobileNumber: command.User.MobileNumber,
	// 	Username:     command.User.Username,
	// 	Password:     command.User.Password,
	// 	IsPrivate:    command.User.IsPrivate,
	// }
	fmt.Println("aa "+command.User.Id)

	reply := events.ChangePrivacyReply{User: command.User}

	switch command.Type {
	case events.ChangePrivacy:
		_, err := handler.userService.ChangeAccountPrivacy(&command.User)
		if err != nil {
			reply.Type = events.PrivacyNotChanged
			break
		}
		reply.Type = events.PrivacyChanged
	case events.RollbackUserPrivacy:
		command.User.IsPrivate = !command.User.IsPrivate
		_, err := handler.userService.ChangeAccountPrivacy(&command.User)
		if err != nil {
			return
		}
		reply.Type = events.UserPrivacyRolledBack
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
