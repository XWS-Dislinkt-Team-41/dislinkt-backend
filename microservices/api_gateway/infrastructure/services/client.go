package services

import (
	"log"

	connections "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/connect_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewConnectClient(address string) connections.ConnectServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Connect service: %v", err)
	}
	return connections.NewConnectServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
