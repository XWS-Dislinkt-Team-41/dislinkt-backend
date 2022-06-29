package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role int

const (
	USER Role = iota
	ADMIN
)

type Method int

const (
	POST Role = iota
	GET
	PUT
	DELETE
)

type UserCredential struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Role     Role               `bson:"role"`
}

type Permission struct {
	Id       primitive.ObjectID `bson:"_id"`
	Role     Role               `bson:"role"`
	Method   Method				`bson:"method"`
	Url      string 			`bson:"url"`
}

type JWTToken struct {
	Token string `json:"token"`
}

func ConvertRoleToString(role Role) string {
	if role == 0 {
		return "USER"
	} else {
		return "ADMIN"
	}
}

func ConvertMethodToString(role Role) string {
	if role == 0 {
		return "POST"
	} else if role == 1{
		return "GET"
	}else if role == 2{
		return "PUT"
	}else {
		return "DELETE"
	}
}