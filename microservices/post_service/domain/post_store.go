package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostStore interface {
	Get(id, post_id primitive.ObjectID) (*Post, error)
	GetAll([]string) ([]*Post, error)
	GetAllFromCollection(id primitive.ObjectID) ([]*Post, error)
	Insert(id primitive.ObjectID, post *Post) (*Post, error)
	InsertComment(id primitive.ObjectID, post_id primitive.ObjectID, comment *Comment) (*Comment, error)
	UpdateLikes(reaction *Reaction) (*Post, error)
	RemoveLike(reaction *Reaction) (*Post, error)
	UpdateDislikes(reaction *Reaction) (*Post, error)
	DeleteAll()
}
