package application

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/message_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageService struct {
	store domain.MessageStore
}

func NewMessageService(store domain.MessageStore) *MessageService {
	return &MessageService{
		store: store,
	}
}

func (service *MessageService) Get(id, connectedId primitive.ObjectID) ([]*domain.Message, error) {
	return service.store.Get(id, connectedId)
}

func (service *MessageService) SendMessage(id, connectedId primitive.ObjectID, message *domain.Message) (*domain.Message, error) {
	return service.store.SendMessage(id, connectedId, message)
}
