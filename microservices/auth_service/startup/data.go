package startup

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/auth_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var credentials = []*domain.UserCredential{
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Username:     "leka",
		Password:     "123",
		Role:   	  domain.ADMIN,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc169"),
		Username:     "pape",
		Password:     "123",
		Role:   	  domain.USER,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc168"),
		Username:     "dare",
		Password:     "123",
		Role:   	  domain.USER,
	},
}

var permissions = []*domain.Permission{
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.USER,
		Method: 	  domain.GET,
		Url:		  "/jobOffer",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}