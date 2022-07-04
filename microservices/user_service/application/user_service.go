package application

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	events "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/saga/make_account_private"
)

type UserService struct {
	store domain.UserStore
	orchestrator    *PrivateAccountOrchestrator
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

func (service *UserService) GetAllPublicUserId() ([]primitive.ObjectID, error) {
	return service.store.GetAllPublicUserId()
}

func (service *UserService) IsPrivate(id primitive.ObjectID) (bool, error) {
	return service.store.IsPrivate(id)
}

func (service *UserService) Register(user *domain.User) (*domain.User, error) {
	return service.store.Insert(user)
}

func (service *UserService) SearchPublic(filter string) ([]*domain.User, error) {
	return service.store.SearchPublic(filter)
}

func (service *UserService) UpdatePersonalInfo(user *domain.User) (*domain.User, error) {
	return service.store.UpdatePersonalInfo(user)
}

func (service *UserService) UpdateCareerInfo(user *domain.User) (*domain.User, error) {
	return service.store.UpdateCareerInfo(user)
}

func (service *UserService) UpdateInterestsInfo(user *domain.User) (*domain.User, error) {
	return service.store.UpdateInterestsInfo(user)
}

func (service *UserService) DeleteById(id primitive.ObjectID) error {
	err := service.store.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) ChangeAccountPrivacy(user *events.UserDetails) (*events.UserDetails, error) {
	var userPrivacy = &domain.User{
		Id: user.Id,
		IsPrivate: user.IsPrivate,
	}
	userInDatabase, err := service.store.UpdateAccountPrivacy(userPrivacy)
	if err != nil {
		return nil,err
	}
	id := userInDatabase.Id.Hex();
	user.id = id;
	err = service.orchestrator.Start(*user)
	if err != nil {
		return nil,err
	}
	return user,nil
}
