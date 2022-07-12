package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatRoom struct {
	Id              primitive.ObjectID `bson:"_id"`
	Name            string             `bson:"name"`
	Image           string             `bson:"image"`
	ParticipantsIds []string           `bson:"participantsIds"`
	Messages        []Message          `bson:"messages"`
}

type Message struct {
	Id       primitive.ObjectID `bson:"_id"`
	Text     string             `bson:"text"`
	SentTime time.Time          `bson:"sentTime"`
	Seen     bool               `bson:"seen"`
}

type UserChatRooms struct {
	Id        primitive.ObjectID `bson:"_id"`
	UserId    string             `bson:"userId"`
	ChatRooms []string
}
