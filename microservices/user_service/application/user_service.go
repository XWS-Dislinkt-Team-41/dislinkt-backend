package application

import (
	"fmt"

	events "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/saga/change_account_privacy"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	store        domain.UserStore
	orchestrator *ChangePrivacyOrchestrator
}

func NewUserService(store domain.UserStore, orchestrator *ChangePrivacyOrchestrator) *UserService {
	return &UserService{
		store:        store,
		orchestrator: orchestrator,
	}
}

func (service *UserService) Get(id primitive.ObjectID) (*domain.User, error) {
	return service.store.Get(id)
}

func (service *UserService) GetPrincipal(principalUsername string) (*domain.User, error) {
	return service.store.GetByUsername(principalUsername)
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

func (service *UserService) Search(filter string) ([]*domain.User, error) {
	return service.store.Search(filter)
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
	id, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return nil, err
	}
	var userPrivacy = &domain.User{
		Id:        id,
		IsPrivate: user.IsPrivate,
	}
	_, err = service.store.UpdateAccountPrivacy(userPrivacy)
	if err != nil {
		return nil, err
	}
	err = service.orchestrator.Start(*user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) RollbackAccountPrivacy(user *events.UserDetails) (*events.UserDetails, error) {
	id, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return nil, err
	}
	var userPrivacy = &domain.User{
		Id:        id,
		IsPrivate: !user.IsPrivate,
	}
	fmt.Println(userPrivacy.IsPrivate)
	_, err = service.store.UpdateAccountPrivacy(userPrivacy)
	if err != nil {
		return nil, err
	}
	return user, nil
}
