package config

import "os"

type Config struct {
	Port                       string
	UserDBHost                 string
	UserDBPort                 string
	NatsHost                   string
	NatsPort                   string
	NatsUser                   string
	NatsPass                   string
	RegisterUserCommandSubject string
	RegisterUserReplySubject   string
}

func NewConfig() *Config {
	return &Config{
		Port:                       os.Getenv("USER_SERVICE_PORT"),
		UserDBHost:                 os.Getenv("USER_DB_HOST"),
		UserDBPort:                 os.Getenv("USER_DB_PORT"),
		NatsHost:                   os.Getenv("NATS_HOST"),
		NatsPort:                   os.Getenv("NATS_PORT"),
		NatsUser:                   os.Getenv("NATS_USER"),
		NatsPass:                   os.Getenv("NATS_PASS"),
		RegisterUserCommandSubject: os.Getenv("REGISTER_USER_COMMAND_SUBJECT"),
		RegisterUserReplySubject:   os.Getenv("REGISTER_USER_REPLY_SUBJECT"),
	}
}
