package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConnectStore interface {
	Register(user Profile) (*Profile, error)
	UpdateUser(user Profile) (*Profile, error)
	IsUserPrivate(userId primitive.ObjectID) (*bool, error)
	Connect(userId, cUserId primitive.ObjectID) (*Connection, error)
	UnConnect(userId, cUserId primitive.ObjectID) error
	GetUserConnections(userId primitive.ObjectID) ([]*Connection, error)
	Invite(userId, cUserId primitive.ObjectID) (*Connection, error)
	AcceptInvitation(userId, cUserId primitive.ObjectID) (*Connection, error)
	DeclineInvitation(userId, cUserId primitive.ObjectID) error
	CancelInvitation(userId, cUserId primitive.ObjectID) error
	GetAllInvitations(userId primitive.ObjectID) ([]*Connection, error)
	GetAllSentInvitations(userId primitive.ObjectID) ([]*Connection, error)
	InitNeo4jDB() error
	GetUserSuggestions(userId primitive.ObjectID) ([]*Profile, error)
	GetRandomUsers(userId primitive.ObjectID) ([]*Profile, error)
	GetRandomUsersWithoutConections(userId primitive.ObjectID) ([]*Profile, error)
}
