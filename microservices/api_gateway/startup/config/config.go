package config

import "os"

type Config struct {
	Port        string
	ConnectHost string
	ConnectPort string
}

func NewConfig() *Config {
	return &Config{
		Port:        os.Getenv("GATEWAY_PORT"),
		ConnectHost: os.Getenv("CONNECT_SERVICE_HOST"),
		ConnectPort: os.Getenv("CONNECT_SERVICE_PORT"),
	}
}
