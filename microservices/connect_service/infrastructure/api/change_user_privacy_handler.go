package api

import (
	saga "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/saga/messaging"
	events "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/saga/change_account_privacy"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/application"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChangePrivacyCommandHandler struct {
	connectService    *application.ConnectService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewChangePrivacyCommandHandler(connectService *application.ConnectService, publisher saga.Publisher, subscriber saga.Subscriber) (*ChangePrivacyCommandHandler, error) {
	o := &ChangePrivacyCommandHandler{
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

func (handler *ChangePrivacyCommandHandler) handle(command *events.ChangePrivacyCommand) {
	id, err := primitive.ObjectIDFromHex(command.User.Id)
	if err != nil {
		return
	}
	user := domain.Profile{
		Id:           id,
		Private:      command.User.IsPrivate,
	}

	reply := events.ChangePrivacyReply{User: command.User}

	switch command.Type {
	case events.ChangePrivacyNode:
		_, err := handler.connectService.UpdateUser(user)
		if err != nil {
			reply.Type = events.PrivacyNodeNotChanged
			break
		}
		reply.Type = events.PrivacyNodeChanged
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
