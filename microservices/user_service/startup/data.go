package startup

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var users = []*domain.User{
	{
		Id:           getObjectId("623b0cc3a34d25d8567f9f82"),
		Firstname:    "Dusan",
		Email:        "leka@gmail.com",
		MobileNumber: "0615656582",
		Gender:       domain.Male,
		BirthDay:     time.Now(),
		Username:     "leka",
		Biography:    "Bio",
		Experience:   "Radio na farmi",
		Education:    "Srednje",
		Skills:       "Programiranje",
		Interests:    "Biciklizam",
		Password:     "123",
	},
	{
		Id:           getObjectId("623b0cc3a34d25d8567f9f83"),
		Firstname:    "Aleksa",
		Email:        "pape@gmail.com",
		MobileNumber: "06586868684",
		Gender:       domain.Male,
		BirthDay:     time.Now(),
		Username:     "pape",
		Biography:    "Neka biografija",
		Experience:   "BIOz",
		Education:    "Fakultet",
		Skills:       "Programiranje",
		Interests:    "Gejminh",
		Password:     "123",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
