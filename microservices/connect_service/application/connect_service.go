package application

import (
	"context"
	"fmt"
	"os"

	notificationService "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/notification_service"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConnectService struct {
	store               domain.ConnectStore
	NotificationService notificationService.NotificationServiceClient
}

func NewConnectService(store domain.ConnectStore) *ConnectService {
	return &ConnectService{
		store:               store,
		NotificationService: NewNotificationClient(fmt.Sprintf("%s:%s", os.Getenv("NOTIFICATION_SERVICE_HOST"), os.Getenv("NOTIFICATION_SERVICE_PORT"))),
	}
}

func (service *ConnectService) Register(profile domain.Profile) (*domain.Profile, error) {
	user, err := service.store.Register(profile)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *ConnectService) UpdateUser(profile domain.Profile) (*domain.Profile, error) {
	user, err := service.store.UpdateUser(profile)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *ConnectService) Connect(userId, cUserId primitive.ObjectID) (*domain.Connection, error) {
	var err error
	var connection *domain.Connection
	isPrivate, err := service.IsProfilePrivate(cUserId)
	if err != nil {
		return nil, err
	}
	if *isPrivate {
		connection, err = service.store.Invite(userId, cUserId)
		if err != nil {
			return nil, err
		}
		var notification notificationService.Notification
		notification.OwnerId = cUserId.Hex()
		notification.ForwardUrl = "requests"
		notification.Text = "sent you a friend request"
		notification.Type = notificationService.Notification_CONNECT
		service.NotificationService.InsertNotification(context.TODO(), &notificationService.InsertNotificationRequest{Notification: &notification})
	} else {
		connection, err = service.store.Connect(userId, cUserId)
		if err != nil {
			return nil, err
		}
		var notification notificationService.Notification
		notification.OwnerId = userId.Hex()
		notification.ForwardUrl = "profile/" + cUserId.Hex()
		notification.Text = "is now your connection"
		notification.Type = notificationService.Notification_CONNECT
		service.NotificationService.InsertNotification(context.TODO(), &notificationService.InsertNotificationRequest{Notification: &notification})
	}
	if err != nil {
		return nil, err
	}
	return connection, nil
}

func (service *ConnectService) UnConnect(userId, cUserId primitive.ObjectID) error {
	err := service.store.UnConnect(userId, cUserId)
	if err != nil {
		return err
	}
	return nil
}

func (service *ConnectService) GetUserConnections(userId primitive.ObjectID) ([]*domain.Connection, error) {
	connections, err := service.store.GetUserConnections(userId)
	if err != nil {
		return nil, err
	}
	return connections, err
}

func (service *ConnectService) AcceptInvitation(userId, cUserId primitive.ObjectID) (*domain.Connection, error) {
	invitation, err := service.store.AcceptInvitation(userId, cUserId)
	if err != nil {
		return nil, err
	}
	return invitation, nil
}

func (service *ConnectService) DeclineInvitation(userId, cUserId primitive.ObjectID) error {
	err := service.store.DeclineInvitation(userId, cUserId)
	if err != nil {
		return err
	}
	return nil
}

func (service *ConnectService) CancelInvitation(userId, cUserId primitive.ObjectID) error {
	err := service.store.CancelInvitation(userId, cUserId)
	if err != nil {
		return err
	}
	return nil
}

func (service *ConnectService) GetAllInvitations(userId primitive.ObjectID) ([]*domain.Connection, error) {
	invitations, err := service.store.GetAllInvitations(userId)
	if err != nil {
		return nil, err
	}
	return invitations, err
}

func (service *ConnectService) GetAllSentInvitations(userId primitive.ObjectID) ([]*domain.Connection, error) {
	invitations, err := service.store.GetAllSentInvitations(userId)
	if err != nil {
		return nil, err
	}
	return invitations, err
}

func (service *ConnectService) IsProfilePrivate(userId primitive.ObjectID) (*bool, error) {
	isUserPrivate, err := service.store.IsUserPrivate(userId)
	if err != nil {
		return nil, err
	}
	return isUserPrivate, err
}

func (service *ConnectService) GetUserSuggestions(userId primitive.ObjectID) ([]*domain.Profile, error) {
	users, err := service.store.GetUserSuggestions(userId)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		users, err = service.store.GetRandomUsers(userId)
		if err != nil {
			return nil, err
		}
	} else if len(users) < 15 {
		randomUsers, err := service.store.GetRandomUsersWithoutConections(userId)
		if err != nil {
			return nil, err
		}
		blockedUsers, err := service.store.GetBlockedUsers(userId)
		if err != nil {
			return nil, err
		}
		for i := 0; i < len(randomUsers) && len(users) < 15; i++ {
			if UserNotSuggested(users, randomUsers[i]) && UserNotBlocked(blockedUsers, randomUsers[i]) {
				users = append(users, randomUsers[i])
			}
		}
	}
	return users, err
}

func UserNotSuggested(suggestedUsers []*domain.Profile, user *domain.Profile) bool {
	for i := 0; i < len(suggestedUsers); i++ {
		if suggestedUsers[i].Id == user.Id {
			return false
		}
	}
	return true
}

func UserNotBlocked(blocks []*domain.Block, user *domain.Profile) bool {
	for i := 0; i < len(blocks); i++ {
		if blocks[i].BUserId == user.Id {
			return false
		}
	}
	return true
}

func (service *ConnectService) Block(userId, bUserId primitive.ObjectID) (*domain.Block, error) {
	block, err := service.store.Block(userId, bUserId)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (service *ConnectService) UnBlock(userId, bUserId primitive.ObjectID) error {
	err := service.store.UnBlock(userId, bUserId)
	if err != nil {
		return err
	}
	return nil
}

func (service *ConnectService) GetBlockedUsers(userId primitive.ObjectID) ([]*domain.Block, error) {
	blocks, err := service.store.GetBlockedUsers(userId)
	if err != nil {
		return nil, err
	}
	return blocks, err
}
