package api

import (
	"context"

	pb "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/notification_service"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/notification_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (handler *NotificationHandler) GetAllNotifications(ctx context.Context, request *pb.GetAllNotificationsRequest) (*pb.GetAllNotificationsResponse, error) {
	userId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}
	notifications, err := handler.service.GetAllNotifications(userId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllNotificationsResponse{
		Notifications: []*pb.Notification{},
	}
	for _, notification := range notifications {
		current := mapNotification(notification)
		response.Notifications = append(response.Notifications, current)
	}
	return response, nil
}

func (handler *NotificationHandler) MarkAllAsSeen(ctx context.Context, request *pb.MarkAllAsSeenRequest) (*pb.EmptyRespones, error) {
	userId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}
	err = handler.service.MarkAllAsSeen(userId)
	if err != nil {
		return nil, err
	}
	response := &pb.EmptyRespones{}
	return response, nil
}

func (handler *NotificationHandler) InsertNotification(ctx context.Context, request *pb.InsertNotificationRequest) (*pb.InsertNotificationResponse, error) {
	notification, err := mapNotificationPb(request.Notification)
	if err != nil {
		return nil, err
	}
	notification, err = handler.service.InsertNotification(notification)
	if err != nil {
		return nil, err
	}
	response := &pb.InsertNotificationResponse{
		Notification: mapNotification(notification),
	}
	return response, nil
}

func (handler *NotificationHandler) GetUserSettings(ctx context.Context, request *pb.GetUserSettingsRequest) (*pb.GetUserSettingsResponse, error) {
	userId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}
	userSettings, err := handler.service.GetUserSettings(userId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetUserSettingsResponse{
		UserSettings: mapUserSettings(userSettings),
	}
	return response, nil
}

func (handler *NotificationHandler) UpdateUserSettings(ctx context.Context, request *pb.UpdateUserSettingsRequest) (*pb.GetUserSettingsResponse, error) {
	userSettings, err := mapUserSettingsPb(request.UserSettings)
	if err != nil {
		return nil, err
	}
	userSettings, err = handler.service.UpdateUserSettings(userSettings)
	if err != nil {
		return nil, err
	}
	response := &pb.GetUserSettingsResponse{
		UserSettings: mapUserSettings(userSettings),
	}
	return response, nil
}
