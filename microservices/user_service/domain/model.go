package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
type GenderEnum int8

const (
	Male GenderEnum = iota
	Female
)

func (status GenderEnum) String() string {
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
	Gender       GenderEnum         `bson:"gender"`
	BirthDay     time.Time          `bson:"birthDay"`
	Username     string             `bson:"username"`
	Biography    string             `bson:"biography"`
	Experience   string             `bson:"experience"`
	Education    string      		`bson:"education"`
	Skills       string             `bson:"skills"`
	Interests    string             `bson:"interests"`
	Password     string             `bson:"password"`
	IsPrivate      bool				`bson:"isPrivate"`
}