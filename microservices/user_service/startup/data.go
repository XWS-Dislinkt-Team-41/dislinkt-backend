package startup

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var users = []*domain.User{
	{
		Id:           getObjectId("626ed920b5d7948d48ffc170"),
		Firstname:    "Dusan",
		Email:        "leka@gmail.com",
		MobileNumber: "0615656582",
		Gender:       domain.Male,
		BirthDay:     time.Now(),
		Username:     "leka",
		Biography:    "Bio",
		Experience:   "Praksa u Excaliburu",
		Education:    "Srednje",
		Skills:       "Programiranje",
		Interests:    "Biciklizam",
		Password:     "123",
		IsPrivate:    false,
	},
	{
		Id:           getObjectId("626ed920b5d7948d48ffc169"),
		Firstname:    "Aleksa",
		Email:        "pape@gmail.com",
		MobileNumber: "06586868684",
		Gender:       domain.Male,
		BirthDay:     time.Now(),
		Username:     "pape",
		Biography:    "Bio",
		Experience:   "Radio u Vega IT",
		Education:    "Fakultet",
		Skills:       "Programiranje",
		Interests:    "Gejming",
		Password:     "123",
		IsPrivate:    true,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
