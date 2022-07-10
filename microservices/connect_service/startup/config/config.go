package config

import "os"

type Config struct {
	Port                       string
	ConnectDBHost              string
	ConnectDBUser              string
	ConnectDBPass              string
	ConnectDBPort              string
	NatsHost                   string
	NatsPort                   string
	NatsUser                   string
	NatsPass                   string
	RegisterUserCommandSubject string
	RegisterUserReplySubject   string
	ChangeAccountPrivacyCommandSubject string
	ChangeAccountPrivacyReplySubject   string
}

func NewConfig() *Config {
	return &Config{
		Port:                       os.Getenv("CONNECT_SERVICE_PORT"),
		ConnectDBHost:              os.Getenv("CONNECT_DB_HOST"),
		ConnectDBUser:              os.Getenv("CONNECT_DB_USER"),
		ConnectDBPass:              os.Getenv("CONNECT_DB_PASS"),
		ConnectDBPort:              os.Getenv("CONNECT_DB_PORT"),
		NatsHost:                   os.Getenv("NATS_HOST"),
		NatsPort:                   os.Getenv("NATS_PORT"),
		NatsUser:                   os.Getenv("NATS_USER"),
		NatsPass:                   os.Getenv("NATS_PASS"),
		RegisterUserCommandSubject: os.Getenv("REGISTER_USER_COMMAND_SUBJECT"),
		RegisterUserReplySubject:   os.Getenv("REGISTER_USER_REPLY_SUBJECT"),
		ChangeAccountPrivacyCommandSubject: os.Getenv("CHANGE_ACCOUNT_PRIVACY_COMMAND_SUBJECT"),
		ChangeAccountPrivacyReplySubject:   os.Getenv("CHANGE_ACCOUNT_PRIVACY_REPLY_SUBJECT"),
	}
}
