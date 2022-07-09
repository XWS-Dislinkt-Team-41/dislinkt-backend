package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type NotificationStore interface {
	DeleteAllSettings()
	UpdateOrInsertSetting(setting *UserSettings) (*UserSettings, error)
	InsertSetting(setting *UserSettings) (*UserSettings, error)
	GetOrInitUserSetting(userId primitive.ObjectID) (*UserSettings, error)

	GetAll() ([]*Notification, error)
	Insert(notification *Notification) (*Notification, error)
	DeleteAllNotifications()
	MarkAsSeen(notificationId primitive.ObjectID) error
}
