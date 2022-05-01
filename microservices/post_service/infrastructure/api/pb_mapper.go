package api

import (
	pb "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/post_service"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/post_service/domain"
)

func mapPost(post *domain.Post) *pb.Post {
	postPb := &pb.Post{
		Id:       post.Id.Hex(),
		Text:     post.Text,
		Link:     post.Link,
		Image:    post.Image,
		OwnerId:  post.OwnerId.Hex(),
		Likes:    post.Likes,
		Dislikes: post.Dislikes,
	}
	for _, comment := range post.Comments {
		postPb.Comments = append(postPb.Comments, &pb.Comment{
			Code: comment.Code,
			Text: comment.Text,
		})
	}
	return postPb
}

func mapNewPost(postPb *pb.Post) *domain.Post {
	post := &domain.Post{
		Comments: make([]domain.Comment, 0),
	}
	for _, commentPb := range postPb.Comments {
		comment := domain.Comment{
			Code: commentPb.Code,
			Text: commentPb.Text,
		}
		post.Comments = append(post.Comments, comment)
	}
	return post
}
