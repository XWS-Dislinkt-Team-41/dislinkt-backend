package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type AuthStore interface {
	Login(user *UserCredential) (*UserCredential, error)
	Register(user *UserCredential) (*UserCredential, error)
	Insert(user *UserCredential) (*UserCredential, error)
	GetByUsername(username string) (*UserCredential, error)
	DeleteAll()
	DeleteById(id primitive.ObjectID) error
}
