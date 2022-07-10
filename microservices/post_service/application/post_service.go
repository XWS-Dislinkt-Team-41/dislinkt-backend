package application

import (
	"context"
	"fmt"
	"os"

	connectService "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/connect_service"
	notificationService "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/notification_service"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostService struct {
	store               domain.PostStore
	NotificationService notificationService.NotificationServiceClient
	ConnectService      connectService.ConnectServiceClient
}

func NewPostService(store domain.PostStore) *PostService {
	return &PostService{
		store:               store,
		NotificationService: NewNotificationClient(fmt.Sprintf("%s:%s", os.Getenv("NOTIFICATION_SERVICE_HOST"), os.Getenv("NOTIFICATION_SERVICE_PORT"))),
		ConnectService:      NewConnectClient(fmt.Sprintf("%s:%s", os.Getenv("CONNECT_SERVICE_HOST"), os.Getenv("CONNECT_SERVICE_PORT"))),
	}
}

func (service *PostService) Get(id, post_id primitive.ObjectID) (*domain.Post, error) {
	return service.store.Get(id, post_id)
}

func (service *PostService) GetAll(postIds []string) ([]*domain.Post, error) {
	return service.store.GetAll(postIds)
}

func (service *PostService) GetAllFromCollection(id primitive.ObjectID) ([]*domain.Post, error) {
	return service.store.GetAllFromCollection(id)
}

func (service *PostService) Insert(id primitive.ObjectID, post *domain.Post) (*domain.Post, error) {
	newPost, err := service.store.Insert(id, post)
	if err != nil {
		return nil, err
	}
	connections, _ := service.ConnectService.GetUserConnections(context.TODO(), &connectService.GetUserConnectionsRequest{UserId: id.Hex()})
	for _, connection := range connections.Connections {
		var notification notificationService.Notification
		notification.OwnerId = connection.CUserId
		notification.ForwardUrl = "post/" + post.Id.Hex()
		notification.Text = "posted on their profile"
		service.NotificationService.InsertNotification(context.TODO(), &notificationService.InsertNotificationRequest{Notification: &notification})
	}

	return newPost, nil
}

func (service *PostService) InsertComment(id primitive.ObjectID, post_id primitive.ObjectID, comment *domain.Comment) (*domain.Comment, error) {
	newComment, err := service.store.InsertComment(id, post_id, comment)
	if err != nil {
		return nil, err
	}
	var notification notificationService.Notification
	notification.OwnerId = id.Hex()
	notification.ForwardUrl = "post/" + post_id.Hex()
	notification.Text = "commented your post"
	service.NotificationService.InsertNotification(context.TODO(), &notificationService.InsertNotificationRequest{Notification: &notification})
	return newComment, nil
}

func (service *PostService) UpdateLikes(reaction *domain.Reaction) (*domain.Post, error) {
	updatedPost, err := service.store.UpdateLikes(reaction)
	if err != nil {
		return nil, err
	}
	return updatedPost, nil
}

func (service *PostService) RemoveLike(reaction *domain.Reaction) (*domain.Post, error) {
	updatedPost, err := service.store.RemoveLike(reaction)
	if err != nil {
		return nil, err
	}
	return updatedPost, nil
}

func (service *PostService) UpdateDislikes(reaction *domain.Reaction) (*domain.Post, error) {
	updatedPost, err := service.store.UpdateDislikes(reaction)
	if err != nil {
		return nil, err
	}
	return updatedPost, nil
}
