package main

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/message_service/startup"
	cfg "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/message_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
