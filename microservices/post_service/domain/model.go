package domain

import "go.mongodb.org/mongo-driver/bson/primitive"


type Comment struct{
	Code string `bson:"_id"`
	Text string `bson:"_id"`
}

type Post struct {
	Id            primitive.ObjectID `bson:"_id"`
	Text          string             `bson:"text"`
	Link          string             `bson:"link"`
	Image         string             `bson:"image"`
	OwnerId       primitive.ObjectID `bson:"owner_Id"`
	Likes         int64              `bson:"likes"`
	Dislikes      int64              `bson:"dislikes"`
	Comments      []Comment          `bson:"comments"`
}
