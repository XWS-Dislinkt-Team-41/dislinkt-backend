package main

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/auth_service/startup"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/auth_service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
