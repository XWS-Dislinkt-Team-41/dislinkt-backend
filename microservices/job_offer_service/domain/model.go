package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobOffer struct {
	Id              primitive.ObjectID `bson:"_id"`
	UserId          primitive.ObjectID `bson:"userId"`
	Position        string             `bson:"position"`
	Seniority		string			   `bson:"seniority"`
	Description     string 			   `bson:"description"`
	Prerequisites   []string           `bson:"prerequisites"`
}