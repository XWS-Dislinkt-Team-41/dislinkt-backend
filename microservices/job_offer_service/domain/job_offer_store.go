package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobOfferStore interface {
	Get(id primitive.ObjectID) (*JobOffer, error)
	GetAll() ([]*JobOffer, error)
	Search(content string) ([]*JobOffer, error)
	Insert(jobOffer *JobOffer) (*JobOffer, error)
	Update(jobOffer *JobOffer) (*JobOffer, error)
	DeleteAll()
}
