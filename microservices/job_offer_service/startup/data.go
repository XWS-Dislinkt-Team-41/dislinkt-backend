package startup

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/job_offer_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jobOffers = []*domain.JobOffer{
	{
		Id:             getObjectId("626ed920b5d7948d48ffc171"),
		UserId:         getObjectId("626ed920b5d7948d48ffc170"),
		Company: 		"Continental",
		Seniority:      "Junior",
		Position:       "Software Engineer",
		Description:    "Dobar",
		Prerequisites:  []string{"Praksa u Excaliburu"},
	},
	{
		Id:             getObjectId("626ed920b5d7948d48ffc172"),
		UserId:         getObjectId("626ed920b5d7948d48ffc170"),
		Company: 		"Levi9",
		Seniority:      "Senior",
		Position:       "Embedded Engineer",
		Description:    "Dobar",
		Prerequisites:  []string{"Praksa u RT-Rk"},
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
