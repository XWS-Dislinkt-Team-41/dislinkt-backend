package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageStore interface {
	Get(id, connectedId primitive.ObjectID) ([]*Message, error)
	SendMessage(id, connectedId primitive.ObjectID, message *Message) (*Message, error)
	DeleteAll()
}
