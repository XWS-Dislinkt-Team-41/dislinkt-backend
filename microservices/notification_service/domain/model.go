package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationType int8

const (
	Connect NotificationType = iota
	Post
	Message
)

type UserSettings struct {
	Id                      primitive.ObjectID `bson:"_id,omitempty"`
	OwnerId                 primitive.ObjectID `bson:"ownerId"`
	PostNotifications       bool               `bson:"postNotifications"`
	ConnectionNotifications bool               `bson:"connectionNotifications"`
	MessageNotifications    bool               `bson:"messageNotifications"`
}

type Notification struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	OwnerId      primitive.ObjectID `bson:"ownerId"`
	ForwardUrl   string             `bson:"forwardUrl"`
	Text         string             `bson:"text"`
	Date         time.Time          `bson:"date"`
	Seen         bool               `bson:"seen"`
	UserFullName string             `bson:"userFullName"`
	Type         NotificationType   `bson:"type"`
}
