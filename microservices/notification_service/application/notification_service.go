package application

import (
	"time"

	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/notification_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationService struct {
	store domain.NotificationStore
}

func NewNotificationService(store domain.NotificationStore) *NotificationService {
	return &NotificationService{
		store: store,
	}
}

func (service *NotificationService) GetAllNotifications(userId primitive.ObjectID) ([]*domain.Notification, error) {
	notifications, err := service.store.GetAll()
	if err != nil {
		return nil, err
	}
	var userNotifications []*domain.Notification
	for _, notification := range notifications {
		if notification.OwnerId.Hex() == userId.Hex() {
			userNotifications = append(userNotifications, notification)
		}
	}
	return userNotifications, nil
}

func (service *NotificationService) MarkAllAsSeen(userId primitive.ObjectID) error {
	notifications, err := service.store.GetAll()
	if err != nil {
		return err
	}
	for _, notification := range notifications {
		if notification.OwnerId.Hex() == userId.Hex() {
			if !notification.Seen {
				service.store.MarkAsSeen(notification.Id)
			}
		}
	}
	return nil
}

func (service *NotificationService) InsertNotification(notification *domain.Notification) (*domain.Notification, error) {
	accepts, err := service.UserAcceptsNotification(notification)
	if err != nil {
		return nil, err
	}
	if *accepts {
		notification.Date = time.Now()
		notification.Seen = false
		notification, err = service.store.Insert(notification)
		if err != nil {
			return nil, err
		}
		return notification, nil
	}
	return notification, nil
}

func (service *NotificationService) GetUserSettings(userId primitive.ObjectID) (*domain.UserSettings, error) {
	settings, err := service.store.GetOrInitUserSetting(userId)
	if err != nil {
		return nil, err
	}
	return settings, nil
}

func (service *NotificationService) UpdateUserSettings(settings *domain.UserSettings) (*domain.UserSettings, error) {
	settings, err := service.store.UpdateOrInsertSetting(settings)
	if err != nil {
		return nil, err
	}
	return settings, nil
}

func (service *NotificationService) UserAcceptsNotification(notification *domain.Notification) (*bool, error) {
	settings, err := service.store.GetOrInitUserSetting(notification.OwnerId)
	if err != nil {
		return nil, err
	}
	if notification.Type == domain.Message {
		return &settings.MessageNotifications, nil
	} else if notification.Type == domain.Connect {
		return &settings.ConnectionNotifications, nil
	} else if notification.Type == domain.Post {
		return &settings.PostNotifications, nil
	}

	return nil, nil
}
