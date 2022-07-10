package api

import (
	pb "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/notification_service"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/notification_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapNotificationPb(notificationPb *pb.Notification) (*domain.Notification, error) {
	userId, err := primitive.ObjectIDFromHex(notificationPb.OwnerId)
	if err != nil {
		return nil, err
	}
	notification := &domain.Notification{
		OwnerId:      userId,
		ForwardUrl:   notificationPb.ForwardUrl,
		Text:         notificationPb.Text,
		Date:         notificationPb.Date.AsTime(),
		Seen:         notificationPb.Seen,
		UserFullName: notificationPb.UserFullName,
	}
	return notification, nil
}

func mapNotification(notification *domain.Notification) *pb.Notification {
	if notification == nil {
		return nil
	}
	notificationPb := &pb.Notification{
		OwnerId:      notification.OwnerId.Hex(),
		ForwardUrl:   notification.ForwardUrl,
		Text:         notification.Text,
		Date:         timestamppb.New(notification.Date),
		Seen:         notification.Seen,
		UserFullName: notification.UserFullName,
	}
	return notificationPb
}

func mapUserSettingsPb(userSettingsPb *pb.UserSettings) (*domain.UserSettings, error) {
	userId, err := primitive.ObjectIDFromHex(userSettingsPb.UserId)
	if err != nil {
		return nil, err
	}
	userSettings := &domain.UserSettings{
		OwnerId:                 userId,
		PostNotifications:       userSettingsPb.PostNotifications,
		ConnectionNotifications: userSettingsPb.ConnectionNotifications,
		MessageNotifications:    userSettingsPb.MessageNotifications,
	}
	return userSettings, nil
}

func mapUserSettings(userSettings *domain.UserSettings) *pb.UserSettings {
	if userSettings == nil {
		return nil
	}
	userSettingsPb := &pb.UserSettings{
		UserId:                  userSettings.OwnerId.Hex(),
		PostNotifications:       userSettings.PostNotifications,
		ConnectionNotifications: userSettings.ConnectionNotifications,
		MessageNotifications:    userSettings.MessageNotifications,
	}
	return userSettingsPb
}
