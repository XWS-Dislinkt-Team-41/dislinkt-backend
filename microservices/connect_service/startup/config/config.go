package config

import "os"

type Config struct {
	Port          string
	ConnectDBHost string
	ConnectDBUser string
	ConnectDBPass string
	ConnectDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:          os.Getenv("CONNECT_SERVICE_PORT"),
		ConnectDBHost: os.Getenv("CONNECT_DB_HOST"),
		ConnectDBUser: os.Getenv("CONNECT_DB_USER"),
		ConnectDBPass: os.Getenv("CONNECT_DB_PASS"),
		ConnectDBPort: os.Getenv("CONNECT_DB_PORT"),
	}
}
