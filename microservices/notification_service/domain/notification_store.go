package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type NotificationStore interface {
	DeleteAllSettings()
	ModifyOrInsertSetting(setting *UserSettings)
	InsertSetting(setting *UserSettings) error
	GetOrInitUserSetting(userId primitive.ObjectID) *UserSettings

	GetAll() ([]*Notification, error)
	Insert(notification *Notification) error
	DeleteAllNotifications()
	MarkAsSeen(notificationId primitive.ObjectID)
}
