package application

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostService struct {
	store domain.PostStore
}

func NewPostService(store domain.PostStore) *PostService {
	return &PostService{
		store: store,
	}
}

func (service *PostService) Get(id primitive.ObjectID, post_id primitive.ObjectID) (*domain.Post, error) {
	return service.store.Get(id, post_id)
}

func (service *PostService) GetAll() ([]*domain.Post, error) {
	return service.store.GetAll()
}

func (service *PostService) GetAllFromCollection(id primitive.ObjectID) ([]*domain.Post, error) {
	return service.store.GetAllFromCollection(id)
}

func (service *PostService) Insert(id primitive.ObjectID, post *domain.Post) (*domain.Post, error) {
	newPost, err := service.store.Insert(post.Id, post)
	if err != nil {
		return nil, err
	}
	return newPost, nil
}

func (service *PostService) InsertComment(id primitive.ObjectID, post_id primitive.ObjectID, comment *domain.Comment) (*domain.Comment, error) {
	newComment, err := service.store.InsertComment(id, post_id, comment)
	if err != nil {
		return nil, err
	}
	return newComment, nil
}
