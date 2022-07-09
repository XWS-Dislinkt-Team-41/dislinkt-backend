package api

import (
	pb "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/notification_service"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/notification_service/application"
)

type NotificationHandler struct {
	pb.UnimplementedNotificationServiceServer
	service *application.NotificationService
}

func NewNotificationHandler(service *application.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		service: service,
	}
}
