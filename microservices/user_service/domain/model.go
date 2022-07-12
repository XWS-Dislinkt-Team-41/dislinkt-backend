package domain

import (
	"time"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service/domain/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Gender int8

const (
	Male Gender = iota
	Female
)

func (status Gender) String() string {
	switch status {
	case Male:
		return "Male"
	case Female:
		return "Female"
	}
	return "Unknown"
}

type User struct {
	Id           primitive.ObjectID `bson:"_id"`
	Firstname    string             `bson:"firstname"`
	Lastname 	 string 			`bson:"lastname"`
	Email        string             `bson:"email"`
	MobileNumber string             `bson:"mobileNumber"`
	Gender       Gender             `bson:"gender"`
	BirthDay     time.Time          `bson:"birthDay"`
	Username     string             `bson:"username"`
	Biography    string             `bson:"biography"`
	Experience   []string           `bson:"experience"`
	Education    enums.Education    `bson:"education"`
	Skills       []string           `bson:"skills"`
	Interests    []string           `bson:"interests"`
	Password     string             `bson:"password"`
	IsPrivate      bool				`bson:"isPrivate"`
}