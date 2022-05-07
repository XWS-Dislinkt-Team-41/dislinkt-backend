package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/api_gateway/domain"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/api_gateway/infrastructure/services"
	conn "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/connect_service"
	post "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/post_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type ConnectPostHandler struct {
	postClientAddress    string
	connectClientAddress string
}

func NewConnectPostHandler(postClientAddress, connectClientAddress string) Handler {
	return &ConnectPostHandler{
		postClientAddress:    postClientAddress,
		connectClientAddress: connectClientAddress,
	}
}

func (handler *ConnectPostHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/user/{id}/connect/post", handler.GetPostsByConnectUsers)
	if err != nil {
		fmt.Println("Panika")
		panic(err)
	}
}

func (handler *ConnectPostHandler) GetPostsByConnectUsers(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	publicPost := &domain.PostsGetAllRequest{}

	err := handler.GetConnectedUsersIds(id, publicPost)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if publicPost.Ids == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		err1 := handler.getAllPosts(publicPost)
		if err1 != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		response, err := json.Marshal(publicPost.Posts)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (handler *ConnectPostHandler) GetConnectedUsersIds(id string, publicPost *domain.PostsGetAllRequest) error {
	connectClient := services.NewConnectClient(handler.connectClientAddress)
	response, err := connectClient.GetUserConnections(context.TODO(), &conn.GetUserConnectionsRequest{UserId: id})
	if err != nil {
		return err
	}
	for _, connection := range response.Connections {
		publicPost.Ids = append(publicPost.Ids, connection.CUser.Id)
	}
	return nil
}

func (handler *ConnectPostHandler) getAllPosts(publicPost *domain.PostsGetAllRequest) error {
	postClient := services.NewPostClient(handler.postClientAddress)
	postCollection, err := postClient.GetAll(context.TODO(), &post.GetAllPublicPostsRequest{PostIds: publicPost.Ids})
	if err != nil {
		return err
	}

	publicPost.Posts = postCollection.Posts
	return nil
}
