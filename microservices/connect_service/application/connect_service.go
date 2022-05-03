package application

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConnectService struct {
	store domain.ConnectStore
}

func NewConnectService(store domain.ConnectStore) *ConnectService {
	return &ConnectService{
		store: store,
	}
}

func (service *ConnectService) Connect(userId, cUserId primitive.ObjectID) (*domain.Connection, error) {
	var err error
	var connection *domain.Connection
	isPrivate, err := service.IsProfilePrivate(cUserId)
	if err != nil {
		return nil, err
	}
	if isPrivate {
		connection, err = service.store.Invite(userId, cUserId)
	} else {
		connection, err = service.store.Connect(userId, cUserId)
	}
	if err != nil {
		return nil, err
	}
	return connection, nil
}

func (service *ConnectService) UnConnect(userId, cUserId primitive.ObjectID) error {
	err := service.store.UnConnect(userId, cUserId)
	if err != nil {
		return err
	}
	return nil
}

func (service *ConnectService) GetUserConnections(userId primitive.ObjectID) ([]*domain.Connection, error) {
	connections, err := service.store.GetUserConnections(userId)
	if err != nil {
		return nil, err
	}
	return connections, err
}

func (service *ConnectService) AcceptInvitation(userId, cUserId primitive.ObjectID) (*domain.Connection, error) {
	invitation, err := service.store.AcceptInvitation(userId, cUserId)
	if err != nil {
		return nil, err
	}
	return invitation, nil
}

func (service *ConnectService) DeclineInvitation(userId, cUserId primitive.ObjectID) error {
	err := service.store.DeclineInvitation(userId, cUserId)
	if err != nil {
		return err
	}
	return nil
}

func (service *ConnectService) CancelInvitation(userId, cUserId primitive.ObjectID) error {
	err := service.store.CancelInvitation(userId, cUserId)
	if err != nil {
		return err
	}
	return nil
}

func (service *ConnectService) GetAllInvitations(userId primitive.ObjectID) ([]*domain.Connection, error) {
	invitations, err := service.store.GetAllInvitations(userId)
	if err != nil {
		return nil, err
	}
	return invitations, err
}

func (service *ConnectService) GetAllSentInvitations(userId primitive.ObjectID) ([]*domain.Connection, error) {
	invitations, err := service.store.GetAllSentInvitations(userId)
	if err != nil {
		return nil, err
	}
	return invitations, err
}

func (service *ConnectService) IsProfilePrivate(cUserId primitive.ObjectID) (bool, error) {
	return false, nil
}
