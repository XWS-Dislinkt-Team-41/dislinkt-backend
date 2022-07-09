package persistence

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE            = "notification"
	COLLECTION          = "notifications"
	SETTINGS_COLLECTION = "user_settings"
)

type NotificationMongoDBStore struct {
	notifications *mongo.Collection
	userSettings  *mongo.Collection
}

func NewNotificationMongoDBStore(client *mongo.Client) domain.NotificationStore {
	notifications := client.Database(DATABASE).Collection(COLLECTION)
	settings := client.Database(DATABASE).Collection(SETTINGS_COLLECTION)
	return &NotificationMongoDBStore{
		notifications: notifications,
		userSettings:  settings,
	}
}
