package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Connection struct {
	UserId  primitive.ObjectID `bson:"userId"`
	CUserId primitive.ObjectID `bson:"cUserId"`
}

type Block struct {
	UserId  primitive.ObjectID `bson:"userId"`
	BUserId primitive.ObjectID `bson:"bUserId"`
}

type Profile struct {
	Id      primitive.ObjectID `bson:"id"`
	Private bool               `bson:"private"`
}
