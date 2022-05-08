package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStore interface {
	Get(id primitive.ObjectID) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByUsername(username string) (*User, error)
	GetAll() ([]*User, error)
	GetAllPublicUserId() ([]primitive.ObjectID, error)
	IsPrivate(id primitive.ObjectID) (bool, error)
	Insert(user *User) (*User, error)
	DeleteAll()
	SearchPublic(filter string) ([]*User, error)
	UpdatePersonalInfo(user *User) (string, error)
	UpdateCareerInfo(user *User) (string, error)
	UpdateInterestsInfo(user *User) (string, error)
}
