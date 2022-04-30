package application

import "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/domain"

type ConnectService struct {
	store domain.ConnectStore
}

func NewConnectService(store domain.ConnectStore) *ConnectService {
	return &ConnectService{
		store: store,
	}
}

func (service *ConnectService) Connect(user, userConnect string) error {
	err := service.store.Connect(user, userConnect)
	if err != nil {
		return err
	}
	return nil
}

func (service *ConnectService) UnConnect(user, userConnect string) error {
	err := service.store.UnConnect(user, userConnect)
	if err != nil {
		return err
	}
	return nil
}

func (service *ConnectService) GetUserConnections(user string) ([]string, error) {
	connections, err := service.store.GetUserConnections(user)
	if err != nil {
		return nil, err
	}
	return connections, err
}
