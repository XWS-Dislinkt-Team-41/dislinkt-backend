package persistence

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "connect"
	COLLECTION = "connect_request"
)

type ConnectRequestMongoDBStore struct {
	requests *mongo.Collection
}

func NewProductMongoDBStore(client *mongo.Client) domain.ConnectRequestStore {
	requests := client.Database(DATABASE).Collection(COLLECTION)
	return &ConnectRequestMongoDBStore{
		requests: requests,
	}
}
