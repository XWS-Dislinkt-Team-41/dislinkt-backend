package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Connection struct {
	User  Profile `bson:"user"`
	CUser Profile `bson:"cUser"`
}

type Profile struct {
	Id      primitive.ObjectID `bson:"id"`
	Private bool               `bson:"private"`
}
