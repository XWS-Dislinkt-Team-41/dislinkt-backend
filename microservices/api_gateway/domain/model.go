package domain

import (
	"time"

	post "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/post_service"
)

type User struct {
	Id           string
	Firstname    string
	Lastname     string
	Email        string
	MobileNumber string
	Gender       string
	BirthDay     time.Time
	Username     string
	Biography    string
	Experience   string
	Education    string
	Skills       string
	Interests    string
	Password     string
	IsPrivate    bool
}

type UserStatusRequest struct {
	Id        string
	IsPrivate bool
	Posts     []*post.Post
}

type PostsGetAllRequest struct {
	Ids   []string
	Posts []*post.Post
}

type Connection struct {
	User  Profile
	CUser Profile
}

type Profile struct {
	Id string
}
