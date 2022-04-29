package application

import "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/domain"

type ConnectionService struct {
	store domain.ConnectionStore
}

func NewConnectionService(store domain.ConnectionStore) *ConnectionService {
	return &ConnectionService{
		store: store,
	}
}

func (service *ConnectionService) Connect(user, userConnect string) error {
	err := service.store.Connect(user, userConnect)
	if err != nil {
		return err
	}
	return nil
}

func (service *ConnectionService) UnConnect() error {
	err := service.store.UnConnect()
	if err != nil {
		return err
	}
	return nil
}
