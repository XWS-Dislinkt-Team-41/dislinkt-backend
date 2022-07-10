package application

import (
	"fmt"
	"log"

	connectService "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/connect_service"
)

func NewConnectClient(address string) connectService.ConnectServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		fmt.Println("Gateway failed to start", "Failed to start")
		log.Fatalf("Failed to start gRPC connect to Catalogue service: %v", err)
	}
	return connectService.NewConnectServiceClient(conn)
}
