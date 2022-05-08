package api

import (
	"time"

	pb "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/post_service"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapPost(post *domain.Post) *pb.Post {
	postPb := &pb.Post{
		Id:         post.Id.Hex(),
		Text:       post.Text,
		Links:      post.Links,
		Image:      post.Image,
		LikedBy:    post.LikedBy,
		DislikedBy: post.DislikedBy,
		OwnerId:    post.OwnerId.Hex(),
		Likes:      post.Likes,
		Dislikes:   post.Dislikes,
	}

	for _, comment := range post.Comments {
		postPb.Comments = append(postPb.Comments, &pb.Comment{
			Code: comment.Code,
			Text: comment.Text,
		})
	}
	postPb.CreatedAt = post.CreatedAt.Time().String()
	return postPb
}

func mapPostRequest(postPb *pb.Post) *domain.Post {
	id, _ := primitive.ObjectIDFromHex(postPb.Id)
	ownerId, _ := primitive.ObjectIDFromHex(postPb.OwnerId)
	Post := &domain.Post{
		Id:         id,
		Text:       postPb.Text,
		Links:      postPb.Links,
		Image:      postPb.Image,
		OwnerId:    ownerId,
		LikedBy:    postPb.LikedBy,
		DislikedBy: postPb.DislikedBy,
		Likes:      postPb.Likes,
		Dislikes:   postPb.Dislikes,
		Comments:   make([]domain.Comment, 0),
	}
	for _, commentPb := range postPb.Comments {
		comment := domain.Comment{
			Code: commentPb.Code,
			Text: commentPb.Text,
		}
		Post.Comments = append(Post.Comments, comment)
	}
	t, _ := time.Parse(time.RFC3339, postPb.CreatedAt)
	Post.CreatedAt = primitive.NewDateTimeFromTime(t)
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
