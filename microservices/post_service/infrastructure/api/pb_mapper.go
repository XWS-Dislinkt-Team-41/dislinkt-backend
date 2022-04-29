package api

import (
	pb "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/post_service"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/post_service/domain"
)

func mapPost(post *domain.Post) *pb.Post {
	postPb := &pb.Post{
		Id:            post.Id.Hex(),
		Name:          post.Name,
		ClothingBrand: post.ClothingBrand,
	}
	return postPb
}
