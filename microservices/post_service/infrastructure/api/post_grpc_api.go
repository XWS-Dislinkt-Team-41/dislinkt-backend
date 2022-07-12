package api

import (
	"context"

	pb "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/post_service"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/post_service/application"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostHandler struct {
	pb.UnimplementedPostServiceServer
	service *application.PostService
}

func NewPostHandler(service *application.PostService) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

func (handler *PostHandler) Get(ctx context.Context, request *pb.GetPostRequest) (*pb.GetResponse, error) {
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	objectPostId, err := primitive.ObjectIDFromHex(request.PostId)
	if err != nil {
		return nil, err
	}
	post, err := handler.service.Get(objectId, objectPostId)
	if err != nil {
		return nil, err
	}
	postPb := mapPost(post)
	response := &pb.GetResponse{
		Post: postPb,
	}
	return response, nil
}

func (handler *PostHandler) GetAll(ctx context.Context, request *pb.GetAllPublicPostsRequest) (*pb.GetAllResponse, error) {
	posts, err := handler.service.GetAll(request.PostIds)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Posts: []*pb.Post{},
	}
	for _, post := range posts {
		current := mapPost(post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (handler *PostHandler) Insert(ctx context.Context, request *pb.NewPostRequest) (*pb.NewPostResponse, error) {
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}

	newPost, err := handler.service.Insert(objectId, mapPostRequest(request.Post))
	response := &pb.NewPostResponse{
		Post: mapPost(newPost),
	}
	return response, err
}

func (handler *PostHandler) InsertComment(ctx context.Context, request *pb.CommentOnPostRequest) (*pb.CommentOnPostResponse, error) {

	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	objectPostId, err := primitive.ObjectIDFromHex(request.PostId)
	if err != nil {
		return nil, err
	}

	newComment, err := handler.service.InsertComment(objectId, objectPostId, mapCommentRequest(request.Comment))
	response := &pb.CommentOnPostResponse{
		Comment: mapComment(newComment),
	}
	return response, err
}

func (handler *PostHandler) InsertReaction(ctx context.Context, request *pb.ReactionOnPostRequest) (*pb.GetResponse, error) {

	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	objectPostId, err := primitive.ObjectIDFromHex(request.PostId)
	if err != nil {
		return nil, err
	}

	objectReactionBy, err := primitive.ObjectIDFromHex(request.Reaction.ReactionBy)
	if err != nil {
		return nil, err
	}
	reaction := &domain.Reaction{
		Id:         objectId,
		PostId:     objectPostId,
		ReactionBy: objectReactionBy,
	}

	var updatedPost *domain.Post
	var err1 error
	if request.Type == "like" {
		updatedPost, err1 = handler.service.UpdateLikes(reaction)
	}
	if request.Type == "dislike" {
		updatedPost, err1 = handler.service.UpdateDislikes(reaction)
	}
	if request.Type == "unlike" {
		updatedPost, err1 = handler.service.RemoveLike(reaction)
	}
	if err1 != nil {
		return nil, err1
	}
	response := &pb.GetResponse{
		Post: mapPost(updatedPost),
	}
	return response, err
}

func (handler *PostHandler) GetAllFromCollection(ctx context.Context, request *pb.GetRequest) (*pb.GetAllResponse, error) {

	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	posts, err := handler.service.GetAllFromCollection(objectId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetAllResponse{
		Posts: []*pb.Post{},
	}
	for _, post := range posts {
		current := mapPost(post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}
