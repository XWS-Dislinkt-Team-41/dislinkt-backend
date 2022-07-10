package persistence

import (
	"context"
	"errors"

	"strings"

	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/job_offer_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "job_offer"
	COLLECTION = "job_offer"
)

type JobOfferMongoDBStore struct {
	jobOffers *mongo.Collection
}

func NewJobOfferMongoDBStore(client *mongo.Client) domain.JobOfferStore {
	jobOffers := client.Database(DATABASE).Collection(COLLECTION)
	return &JobOfferMongoDBStore{
		jobOffers: jobOffers,
	}
}

func (store *JobOfferMongoDBStore) Get(id primitive.ObjectID) (*domain.JobOffer, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *JobOfferMongoDBStore) GetAll() ([]*domain.JobOffer, error) {
	filter := bson.D{}
	return store.filter(filter)
}

func (store *JobOfferMongoDBStore) Search(filter string) ([]*domain.JobOffer, error) {
	var foundJobOffers []*domain.JobOffer

	filter = strings.TrimSpace(filter)
	splitSearch := strings.Split(filter, " ")

	for _, splitSearchpart := range splitSearch {

		//position
		filtereds, err := store.jobOffers.Find(context.TODO(), bson.M{"position": primitive.Regex{Pattern: splitSearchpart, Options: "i"}})
		if err != nil {
			return nil, err
		}
		var jobOffersPosition []*domain.JobOffer
		if err = filtereds.All(context.TODO(), &jobOffersPosition); err != nil {
			return nil, err
		}
		for _, jobOfferOneSlice := range jobOffersPosition {
			foundJobOffers = AppendIfMissing(foundJobOffers, jobOfferOneSlice)
		}

		//description
		filtereds, err = store.jobOffers.Find(context.TODO(), bson.M{"description": primitive.Regex{Pattern: splitSearchpart, Options: "i"}})
		if err != nil {
			return nil, err
		}
		var jobOffersDescription []*domain.JobOffer
		if err = filtereds.All(context.TODO(), &jobOffersDescription); err != nil {
			return nil, err
		}
		for _, jobOfferOneSlice := range jobOffersDescription {
			foundJobOffers = AppendIfMissing(foundJobOffers, jobOfferOneSlice)
		}

		//company
		filtereds, err = store.jobOffers.Find(context.TODO(), bson.M{"company": primitive.Regex{Pattern: splitSearchpart, Options: "i"}})
		if err != nil {
			return nil, err
		}
		var jobOffersCompany []*domain.JobOffer
		if err = filtereds.All(context.TODO(), &jobOffersCompany); err != nil {
			return nil, err
		}
		for _, jobOfferOneSlice := range jobOffersCompany {
			foundJobOffers = AppendIfMissing(foundJobOffers, jobOfferOneSlice)
		}

		//prerequisites
		filtereds, err = store.jobOffers.Find(context.TODO(), bson.M{"prerequisites": primitive.Regex{Pattern: splitSearchpart, Options: "i"}})
		if err != nil {
			return nil, err
		}
		var jobOffersPrerequisites []*domain.JobOffer
		if err = filtereds.All(context.TODO(), &jobOffersPrerequisites); err != nil {
			return nil, err
		}
		for _, jobOfferOneSlice := range jobOffersPrerequisites {
			foundJobOffers = AppendIfMissing(foundJobOffers, jobOfferOneSlice)
		}
	}
	return foundJobOffers, nil
}

func AppendIfMissing(slice []*domain.JobOffer, i *domain.JobOffer) []*domain.JobOffer {
	for _, ele := range slice {
		if ele.Id == i.Id {
			return slice
		}
	}
	return append(slice, i)
}

func (store *JobOfferMongoDBStore) Insert(jobOffer *domain.JobOffer) (*domain.JobOffer, error) {
	filter := bson.M{"id": jobOffer.Id}
	jobOfferInDatabase, _ := store.filterOneRegister(filter)
	if jobOfferInDatabase != nil {
		return nil, errors.New("Job Offer with the same id already exists.")
	}
	jobOffer.Id = primitive.NewObjectID()
	_, err := store.jobOffers.InsertOne(context.TODO(), jobOffer)
	if err != nil {
		return nil, errors.New("Create error.")
	}

	return jobOffer, nil
}

func (store *JobOfferMongoDBStore) DeleteAll() {
	store.jobOffers.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *JobOfferMongoDBStore) filter(filter interface{}) ([]*domain.JobOffer, error) {
	cursor, err := store.jobOffers.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *JobOfferMongoDBStore) Update(jobOffer *domain.JobOffer) (*domain.JobOffer, error) {
	jobOfferInDatabase, err := store.Get(jobOffer.Id)
	if jobOfferInDatabase == nil {
		return nil, err
	}
	jobOfferInDatabase.Description = jobOffer.Description
	jobOfferInDatabase.Position = jobOffer.Position
	jobOfferInDatabase.Prerequisites = jobOffer.Prerequisites
	filter := bson.M{"_id": jobOfferInDatabase.Id}
	update := bson.M{
		"$set": jobOfferInDatabase,
	}
	_, err = store.jobOffers.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return jobOfferInDatabase, nil
}

func (store *JobOfferMongoDBStore) filterOne(filter interface{}) (JobOffer *domain.JobOffer, err error) {
	result := store.jobOffers.FindOne(context.TODO(), filter)
	err = result.Decode(&JobOffer)
	return
}

func decode(cursor *mongo.Cursor) (jobOffers []*domain.JobOffer, err error) {
	for cursor.Next(context.TODO()) {
		var JobOffer domain.JobOffer
		err = cursor.Decode(&JobOffer)
		if err != nil {
			return
		}
		jobOffers = append(jobOffers, &JobOffer)
	}
	err = cursor.Err()
	return
}

func (store *JobOfferMongoDBStore) filterOneRegister(filter interface{}) (jobOffer *domain.JobOffer, err error) {
	result := store.jobOffers.FindOne(context.TODO(), filter)
	err = result.Decode(&jobOffer)
	if err != nil {
		return nil, nil
	}
	return
}
