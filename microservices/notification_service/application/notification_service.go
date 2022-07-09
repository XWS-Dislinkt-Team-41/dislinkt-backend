package application

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/notification_service/domain"
)

type NotificationService struct {
	store domain.NotificationStore
}

func NewNotificationService(store domain.NotificationStore) *NotificationService {
	return &NotificationService{
		store: store,
	}
}
