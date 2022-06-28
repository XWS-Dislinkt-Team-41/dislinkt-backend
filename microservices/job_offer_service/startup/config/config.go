package config

import "os"

type Config struct {
	Port           string
	JobOfferDBHost string
	JobOfferDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:           os.Getenv("JOB_OFFER_SERVICE_PORT"),
		JobOfferDBHost: os.Getenv("JOB_OFFER_DB_HOST"),
		JobOfferDBPort: os.Getenv("JOB_OFFER_DB_PORT"),
	}
}
