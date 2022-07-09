package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type NotificationStore interface {
	DeleteAllSettings()
	UpdateOrInsertSetting(setting *UserSettings)
	InsertSetting(setting *UserSettings) (*UserSettings, error)
	GetOrInitUserSetting(userId primitive.ObjectID) *UserSettings

	GetAll() ([]*Notification, error)
	Insert(notification *Notification) (*Notification, error)
	DeleteAllNotifications()
	MarkAsSeen(notificationId primitive.ObjectID)
}
