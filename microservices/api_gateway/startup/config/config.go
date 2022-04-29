package config

import "os"

type Config struct {
	Port     string
	UserHost string
	UserPort string
	PostHost string
	PostPort string
}

func NewConfig() *Config {
	return &Config{
		Port:     os.Getenv("GATEWAY_PORT"),
		UserHost:  os.Getenv("USER_SERVICE_HOST"),
		UserPort:  os.Getenv("USER_SERVICE_HOST"),
		PostHost: os.Getenv("POST_SERVICE_HOST"),
		PostPort: os.Getenv("POST_SERVICE_PORT"),
	}
}