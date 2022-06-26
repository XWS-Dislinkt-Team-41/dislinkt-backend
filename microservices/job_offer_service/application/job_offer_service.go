package application

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/job_offer_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobOfferService struct {
	store domain.JobOfferStore
}

func NewJobOfferService(store domain.JobOfferStore) *JobOfferService {
	return &JobOfferService{
		store: store,
	}
}

func (service *JobOfferService) Get(id primitive.ObjectID) (*domain.JobOffer, error) {
	return service.store.Get(id)
}

func (service *JobOfferService) GetAll() ([]*domain.JobOffer, error) {
	return service.store.GetAll()
}

func (service *JobOfferService) Search(filter string) ([]*domain.JobOffer, error) {
	return service.store.Search(filter)
}

func (service *JobOfferService) Update(jobOffer *domain.JobOffer) (*domain.JobOffer, error) {
	return service.store.Update(jobOffer)
}

func (service *JobOfferService) Insert(jobOffer *domain.JobOffer) (*domain.JobOffer, error) {
	return service.store.Insert(jobOffer)
}
