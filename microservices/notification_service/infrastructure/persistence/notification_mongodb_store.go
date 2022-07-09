package persistence

import (
	"context"

	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/notification_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (store *NotificationMongoDBStore) DeleteAllNotifications() {
	store.notifications.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *NotificationMongoDBStore) DeleteAllSettings() {
	store.userSettings.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *NotificationMongoDBStore) GetAll() ([]*domain.Notification, error) {
	filter := bson.D{}
	return store.filter(filter)
}

func (store *NotificationMongoDBStore) GetOrInitUserSetting(userId primitive.ObjectID) *domain.UserSettings {
	settingsFilter := bson.M{"ownerId": userId}
	settings, _ := store.filterOneSettings(settingsFilter)
	if settings == nil {
		newSettings := domain.UserSettings{
			OwnerId:                 userId,
			PostNotifications:       true,
			ConnectionNotifications: true,
			MessageNotifications:    true,
		}

		settings, err := store.InsertSetting(&newSettings)
		if err != nil {
			return nil
		}
		return settings
	} else {
		return settings
	}
}

func (store *NotificationMongoDBStore) Insert(notification *domain.Notification) (*domain.Notification, error) {
	notification.Id = primitive.NewObjectID()
	_, err := store.notifications.InsertOne(context.TODO(), notification)
	if err != nil {
		return nil, err
	}
	return notification, nil
}

func (store *NotificationMongoDBStore) InsertSetting(setting *domain.UserSettings) (*domain.UserSettings, error) {
	setting.Id = primitive.NewObjectID()
	_, err := store.userSettings.InsertOne(context.TODO(), setting)
	if err != nil {
		return nil, err
	}
	return setting, nil
}

func (store *NotificationMongoDBStore) MarkAsSeen(notificationId primitive.ObjectID) {
	filter := bson.M{"_id": notificationId}
	updatedNotification := bson.M{"$set": bson.M{
		"seen": true,
	}}
	store.notifications.UpdateOne(context.TODO(), filter, updatedNotification)
}

func (store *NotificationMongoDBStore) UpdateOrInsertSetting(setting *domain.UserSettings) {
	settingOld := store.GetOrInitUserSetting(setting.OwnerId)
	filter := bson.M{"_id": settingOld.Id}
	updatedSetting := bson.M{"$set": bson.M{
		"postNotifications":       setting.PostNotifications,
		"connectionNotifications": setting.ConnectionNotifications,
		"messageNotifications":    setting.MessageNotifications,
	}}
	store.userSettings.UpdateOne(context.TODO(), filter, updatedSetting)
}

func (store *NotificationMongoDBStore) filter(filter interface{}) ([]*domain.Notification, error) {
	cursor, err := store.notifications.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *NotificationMongoDBStore) filterOne(filter interface{}) (notification *domain.Notification, err error) {
	result := store.notifications.FindOne(context.TODO(), filter)
	err = result.Decode(&notification)
	return
}

func (store *NotificationMongoDBStore) filterOneSettings(filter interface{}) (settings *domain.UserSettings, err error) {
	result := store.userSettings.FindOne(context.TODO(), filter)
	err = result.Decode(&settings)
	return
}

func decode(cursor *mongo.Cursor) (notifications []*domain.Notification, err error) {
	for cursor.Next(context.TODO()) {
		var notification domain.Notification
		err = cursor.Decode(&notification)
		if err != nil {
			return
		}
		notifications = append(notifications, &notification)
	}
	err = cursor.Err()
	return
}
