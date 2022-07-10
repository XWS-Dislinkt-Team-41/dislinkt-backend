package application

import (
	"crypto/tls"
	"fmt"
	"log"

	notificationService "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/notification_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func getConnection(address string) (*grpc.ClientConn, error) {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	return grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewTLS(config)))
}

func NewNotificationClient(address string) notificationService.NotificationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		fmt.Println("Gateway failed to start", "Failed to start")
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return notificationService.NewNotificationServiceClient(conn)
}
