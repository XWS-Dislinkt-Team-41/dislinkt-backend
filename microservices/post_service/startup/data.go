package startup

import (
	"time"

	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var posts = []domain.Post{
	{
		Id: getObjectId("623b0cc3a34d25d8567f9f82"),
		OwnerId: getObjectId("626ed920b5d7948d48ffc170"),
		Text: "Praksa u LEVI9",
		CreatedAt : primitive.NewDateTimeFromTime(time.Now().AddDate(0,0,-2)),
	},
	{
		Id: getObjectId("623b0cc3a34d25d8567f9f83"),
		OwnerId: getObjectId("626ed920b5d7948d48ffc170"),
		Text: "Zaposlio se u Microsoftu",
		CreatedAt : primitive.NewDateTimeFromTime(time.Now().AddDate(0,0,-5)),
	},
	{
		Id: getObjectId("623b0cc3a34d25d8567f9f83"),
		OwnerId: getObjectId("626ed920b5d7948d48ffc170"),
		Text: "Zavrsio projekat",
		CreatedAt : primitive.NewDateTimeFromTime(time.Now().AddDate(0,0,-20)),
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
