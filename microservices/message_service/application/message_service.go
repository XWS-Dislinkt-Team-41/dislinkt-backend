package application

import (
	"context"
	"fmt"
	"os"

	notificationService "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/notification_service"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/message_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageService struct {
	store               domain.MessageStore
	NotificationService notificationService.NotificationServiceClient
}

func NewMessageService(store domain.MessageStore) *MessageService {
	return &MessageService{
		store:               store,
		NotificationService: NewNotificationClient(fmt.Sprintf("%s:%s", os.Getenv("NOTIFICATION_SERVICE_HOST"), os.Getenv("NOTIFICATION_SERVICE_PORT"))),
	}
}

func (service *MessageService) Get(id, connectedId primitive.ObjectID) ([]*domain.Message, error) {
	return service.store.Get(id, connectedId)
}

func (service *MessageService) SendMessage(id, connectedId primitive.ObjectID, message *domain.Message) (*domain.Message, error) {
	message, err := service.store.SendMessage(id, connectedId, message)
	if err != nil {
		return nil, err
	}
	var notification notificationService.Notification
	notification.OwnerId = connectedId.Hex()
	notification.ForwardUrl = "chat"
	notification.Text = "sent you a message"
	notification.Type = notificationService.Notification_MESSAGE
	service.NotificationService.InsertNotification(context.TODO(), &notificationService.InsertNotificationRequest{Notification: &notification})
	return message, nil
}
