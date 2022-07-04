package register_user

type UserDetails struct {
	Id           string
	Username     string
	Password     string
	IsPrivate    bool
	Firstname    string
	Lastname     string
	Email        string
	MobileNumber string
	Role         Role
}

type Role int

const (
	USER Role = iota
	ADMIN
)

type RegisterUserCommandType int8

const (
	RegisterUser RegisterUserCommandType = iota
	RollbackRegisterUser
	RegisterUserNode
	RollbackUserCredential
	UnknownCommand
)

type RegisterUserCommand struct {
	User UserDetails
	Type RegisterUserCommandType
}

type RegisterUserReplyType int8

const (
	UserCredentialRolledBack RegisterUserReplyType = iota
	UserRegistered
	UserNotRegistered
	UserRolledBack
	UserNodeRegistered
	UserNodeNotRegistered
	UnknownReply
)

type RegisterUserReply struct {
	User UserDetails
	Type RegisterUserReplyType
}
