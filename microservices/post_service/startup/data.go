package startup

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var posts = []*domain.Post{
	{
		Id:            getObjectId("623b0cc3a34d25d8567f9f82"),
		Name:          "name",
		ClothingBrand: "brand",
	},
	{
		Id:            getObjectId("623b0cc3a34d25d8567f9f83"),
		Name:          "name2",
		ClothingBrand: "brand2",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
