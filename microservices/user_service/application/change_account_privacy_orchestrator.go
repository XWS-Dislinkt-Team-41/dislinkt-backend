package application

import (
	"fmt"

	events "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/saga/change_account_privacy"
	saga "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/saga/messaging"
)

type ChangePrivacyOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewChangePrivacyOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*ChangePrivacyOrchestrator, error) {
	o := &ChangePrivacyOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *ChangePrivacyOrchestrator) Start(userDetails events.UserDetails) error {
	event := &events.ChangePrivacyCommand{
		Type: events.ChangePrivacyNode,
		User: userDetails,
	}
	return o.commandPublisher.Publish(event)
}

func (o *ChangePrivacyOrchestrator) handle(reply *events.ChangePrivacyReply) {
	command := events.ChangePrivacyCommand{User: reply.User}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *ChangePrivacyOrchestrator) nextCommandType(reply events.ChangePrivacyReplyType) events.ChangePrivacyCommandType {
	fmt.Println(reply)
	switch reply {
	case events.PrivacyNodeNotChanged:
		return events.RollbackUserPrivacy
	default:
		return events.UnknownCommand
	}
}
