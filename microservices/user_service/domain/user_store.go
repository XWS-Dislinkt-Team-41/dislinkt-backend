package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStore interface {
	Get(id primitive.ObjectID) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByUsername(username string) (*User, error)
	GetAll() ([]*User, error)
	IsPrivate(id primitive.ObjectID) (bool, error)
	Insert(user *User) (string, error)
	DeleteAll()
	SearchPublic(username string, name string) ([]*User, error)
	UpdatePersonalInfo(user *User) (string, error)
	UpdateCareerInfo(user *User) (string, error)
	UpdateInterestsInfo(user *User) (string, error)
}
