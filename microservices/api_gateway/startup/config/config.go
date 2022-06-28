package config

import "os"

type Config struct {
	Port           string
	UserHost       string
	UserPort       string
	PostHost       string
	PostPort       string
	ConnectHost    string
	ConnectPort    string
	AuthHost       string
	AuthPort       string
	JobOfferHost   string
	JobOfferPort   string
}

func NewConfig() *Config {
	return &Config{
		Port:            os.Getenv("GATEWAY_PORT"),
		UserHost:        os.Getenv("USER_SERVICE_HOST"),
		UserPort:        os.Getenv("USER_SERVICE_PORT"),
		PostHost:        os.Getenv("POST_SERVICE_HOST"),
		PostPort:        os.Getenv("POST_SERVICE_PORT"),
		ConnectHost:     os.Getenv("CONNECT_SERVICE_HOST"),
		ConnectPort:     os.Getenv("CONNECT_SERVICE_PORT"),
		AuthHost:        os.Getenv("AUTH_SERVICE_HOST"),
		AuthPort:        os.Getenv("AUTH_SERVICE_PORT"),
		JobOfferHost:    os.Getenv("JOB_OFFER_SERVICE_HOST"),
		JobOfferPort:    os.Getenv("JOB_OFFER_SERVICE_PORT"),
	}
}
