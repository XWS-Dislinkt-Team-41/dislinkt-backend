package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role int

const (
	USER Role = iota
	ADMIN
)

type Method int

const (
	POST Method = iota
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

func ConvertStringToRole(role string) Role {
	if role == "USER" {
		return USER
	} else {
		return ADMIN
	}
}

func ConvertMethodToString(method Method) string {
	if method == 0 {
		return "POST"
	} else if method == 1{
		return "GET"
	}else if method == 2{
		return "PUT"
	}else {
		return "DELETE"
	}
}

func ConvertStringToMethod(method string) Method {
	if method ==  "POST"{
		return POST
	} else if method == "GET"{
		return GET
	}else if method == "PUT"{
		return PUT
	}else {
		return DELETE
	}
}