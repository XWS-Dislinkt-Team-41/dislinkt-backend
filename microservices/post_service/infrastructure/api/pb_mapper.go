package api

import (
	pb "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/post_service"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func mapPostRequest(postPb *pb.Post) *domain.Post {
	id, _ := primitive.ObjectIDFromHex(postPb.Id)
	ownerId, _ := primitive.ObjectIDFromHex(postPb.OwnerId)
	Post := &domain.Post{
		Id:       id,
		Text:     postPb.Text,
		Link:     postPb.Link,
		Image:    postPb.Image,
		OwnerId:  ownerId,
		Likes:    postPb.Likes,
		Dislikes: postPb.Dislikes,
		Comments: make([]domain.Comment, 0),
	}
	for _, commentPb := range postPb.Comments {
		comment := domain.Comment{
			Code: commentPb.Code,
			Text: commentPb.Text,
		}
		Post.Comments = append(Post.Comments, comment)
	}
	return Post
}

func mapCommentRequest(commentPb *pb.Comment) *domain.Comment {
	Comment := &domain.Comment{
		Code: commentPb.Code,
		Text: commentPb.Text,
	}
	return Comment
}

func mapComment(commentPb *domain.Comment) *pb.Comment {
	Comment := &pb.Comment{
		Code: commentPb.Code,
		Text: commentPb.Text,
	}
	return Comment
}
