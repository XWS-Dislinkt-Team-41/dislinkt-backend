package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ConnectStore interface {
	Connect(userId, cUserId primitive.ObjectID) (*Connection, error)
	UnConnect(userId, cUserId primitive.ObjectID) error
	GetUserConnections(userId primitive.ObjectID) ([]*Connection, error)
	Invite(userId, cUserId primitive.ObjectID) (*Connection, error)
	AcceptInvitation(userId, cUserId primitive.ObjectID) (*Connection, error)
	DeclineInvitation(userId, cUserId primitive.ObjectID) error
	GetAllInvitations(userId primitive.ObjectID) ([]*Connection, error)
	GetAllSentInvitations(userId primitive.ObjectID) ([]*Connection, error)
}
