package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	Code string `bson:"code"`
	Text string `bson:"text"`
}

type Post struct {
	Id         primitive.ObjectID `bson:"_id"`
	Text       string             `bson:"text"`
	Link       string             `bson:"link"`
	Image      string             `bson:"image"`
	OwnerId    primitive.ObjectID `bson:"owner_Id"`
	LikedBy    []string           `bson:"liked_by"`
	DislikedBy []string           `bson:"disliked_by"`
	Likes      int64              `bson:"likes"`
	Dislikes   int64              `bson:"dislikes"`
	Comments   []Comment          `bson:"comments"`
}

type NewPostRequest struct {
	Id   primitive.ObjectID `bson:"_id"`
	Post Post               `bson:"post"`
}

type CommentOnPostRequest struct {
	Id      primitive.ObjectID `bson:"_id"`
	PostID  primitive.ObjectID `bson:"post_id"`
	Comment Comment            `bson:"comment"`
}

type Reaction struct {
	Id         primitive.ObjectID `bson:"_id"`
	PostId     primitive.ObjectID `bson:"post_id"`
	ReactionBy primitive.ObjectID `bson:"reaction"`
}

type ErrorExists struct {
	err error
}
