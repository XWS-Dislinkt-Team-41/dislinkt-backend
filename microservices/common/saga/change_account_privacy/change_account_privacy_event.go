package change_privacy

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

type ChangePrivacyCommandType int8

const (
	RollbackUserPrivacy	ChangePrivacyCommandType = iota
	ChangePrivacyNode
	UnknownCommand
)

type ChangePrivacyCommand struct {
	User UserDetails
	Type ChangePrivacyCommandType
}

type ChangePrivacyReplyType int8

const (
	UserPrivacyRolledBack ChangePrivacyReplyType = iota
	PrivacyNodeChanged
	PrivacyNodeNotChanged
	UnknownReply
)

type ChangePrivacyReply struct {
	User UserDetails
	Type ChangePrivacyReplyType
}
