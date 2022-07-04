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
		Role:		  domain.ADMIN,
		Method: 	  domain.GET,
		Url:		  "/user",
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.USER,
		Method: 	  domain.POST,
		Url:		  `\/user\/[0-9a-f]{24}\/post`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.USER,
		Method: 	  domain.GET,
		Url:		  `\/user\/[0-9a-f]{24}\/public`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.USER,
		Method: 	  domain.PUT,
		Url:		  `\/user\/[0-9a-f]{24}\/post\/[0-9a-f]{24}\/reaction\/like`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.USER,
		Method: 	  domain.GET,
		Url:		  `\/user\/[0-9a-f]{24}\/connect\/post`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.USER,
		Method: 	  domain.PUT,
		Url:		  `\/user\/personal`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.USER,
		Method: 	  domain.PUT,
		Url:		  `\/user\/career`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.USER,
		Method: 	  domain.PUT,
		Url:		  `\/user\/interests`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.USER,
		Method: 	  domain.POST,
		Url:		  `\/user\/[0-9a-f]{24}\/connect`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.USER,
		Method: 	  domain.DELETE,
		Url:		  `\/user\/[0-9a-f]{24}\/connect\/[0-9a-f]{24}`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.USER,
		Method: 	  domain.GET,
		Url:		  `\/user\/[0-9a-f]{24}\/connect`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.USER,
		Method: 	  domain.GET,
		Url:		  `\/user\/[0-9a-f]{24}\/connect\/suggestions`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.USER,
		Method: 	  domain.PUT,
		Url:		  `\/user\/[0-9a-f]{24}\/connect\/invitation\/[0-9a-f]{24}`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.USER,
		Method: 	  domain.DELETE,
		Url:		  `\/user\/[0-9a-f]{24}\/connect\/invitation\/[0-9a-f]{24}`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.USER,
		Method: 	  domain.DELETE,
		Url:		   `\/user\/[0-9a-f]{24}\/connect\/invitation\/[0-9a-f]{24}\/cancel`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.USER,
		Method: 	  domain.GET,
		Url:		   `\/user\/[0-9a-f]{24}\/connect\/invitation`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.USER,
		Method: 	  domain.GET,
		Url:		   `\/user\/[0-9a-f]{24}\/connect\/invitation\/sent`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.USER,
		Method: 	  domain.GET,
		Url:		  `\/jobOffer`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.AGENT,
		Method: 	  domain.GET,
		Url:		  `\/jobOffer`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.USER,
		Method: 	  domain.POST,
		Url:		  `\/jobOffer`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.AGENT,
		Method: 	  domain.POST,
		Url:		  `\/jobOffer`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.USER,
		Method: 	  domain.PUT,
		Url:		  `\/jobOffer`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.AGENT,
		Method: 	  domain.PUT,
		Url:		  `\/jobOffer`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.USER,
		Method: 	  domain.POST,
		Url:		  `\/auth\/connectAgent`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.AGENT,
		Method: 	  domain.POST,
		Url:		  `\/auth\/connectAgent`,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Role:		  domain.USER,
		Method: 	  domain.PUT,
		Url:		  `\/user\/privacy`,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}