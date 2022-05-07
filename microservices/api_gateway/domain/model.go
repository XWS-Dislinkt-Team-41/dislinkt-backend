package domain

import (
	"time"

	auth "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/auth_service"
	post "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/post_service"
	user "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/user_service"
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

type RegisterRequest struct {
	UserCredential auth.UserCredential
	User           user.User
}
