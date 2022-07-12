package startup

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var users = []*domain.Profile{
	{
		Id:      getObjectId("626ed920b5d7948d48ffc170"),
		Private: false,
	},
	{
		Id:      getObjectId("626ed920b5d7948d48ffc169"),
		Private: true,
	},
	{
		Id:      getObjectId("626ed920b5d7948d48ffc168"),
		Private: true,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
