package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostStore interface {
	Get(id, post_id primitive.ObjectID) (*Post, error)
	GetAll() ([]*Post, error)
	GetAllFromCollection(id primitive.ObjectID) ([]*Post, error)
	Insert(id primitive.ObjectID, post *Post) (*Post, error)
	InsertComment(id primitive.ObjectID, post_id primitive.ObjectID, comment *Comment) (*Comment, error)
	UpdateLikes(id primitive.ObjectID, post_id primitive.ObjectID) (*Post, error)
	UpdateDislikes(id primitive.ObjectID, post_id primitive.ObjectID) (*Post, error)
	DeleteAll()
}
