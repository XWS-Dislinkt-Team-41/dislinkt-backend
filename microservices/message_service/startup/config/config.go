package config

import "os"

type Config struct {
	Port          string
	MessageDBHost string
	MessageDBPort string
	NatsHost      string
	NatsPort      string
	NatsMessage   string
	NatsPass      string
}

func NewConfig() *Config {
	return &Config{
		Port:          os.Getenv("MESSAGE_SERVICE_PORT"),
		MessageDBHost: os.Getenv("MESSAGE_DB_HOST"),
		MessageDBPort: os.Getenv("MESSAGE_DB_PORT"),
		NatsHost:      os.Getenv("NATS_HOST"),
		NatsPort:      os.Getenv("NATS_PORT"),
		NatsMessage:   os.Getenv("NATS_USER"),
		NatsPass:      os.Getenv("NATS_PASS"),
	}
}
