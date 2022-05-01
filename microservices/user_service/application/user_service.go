package application

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	store domain.UserStore
}

func NewUserService(store domain.UserStore) *UserService {
	return &UserService{
		store: store,
	}
}

func (service *UserService) Get(id primitive.ObjectID) (*domain.User, error) {
	return service.store.Get(id)
}

func (service *UserService) GetAll() ([]*domain.User, error) {
	return service.store.GetAll()
}

func (service *UserService) Register(user *domain.User) (string, error) {
	message, err := service.store.Insert(user)
	return message, err
}

func (service *UserService) SearchPublic(username string, name string) ([]*domain.User, error) {
	message, err := service.store.SearchPublic(username, name)
	return message, err
}

func (service *UserService) UpdatePersonalInfo(user *domain.User) (string, error) {
	message, err := service.store.UpdatePersonalInfo(user)
	return message, err
}

func (service *UserService) UpdateCareerInfo(user *domain.User) (string, error) {
	message, err := service.store.UpdateCareerInfo(user)
	return message, err
}

func (service *UserService) UpdateInterestsInfo(user *domain.User) (string, error) {
	message, err := service.store.UpdateInterestsInfo(user)
	return message, err
}
