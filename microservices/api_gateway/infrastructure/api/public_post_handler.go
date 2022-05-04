package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/api_gateway/domain"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/api_gateway/infrastructure/services"
	post "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/post_service"
	user "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type PublicPostHandler struct {
	postClientAddress string
	userClientAddress string
}

func NewPublicPostHandler(userClientAddress, postClientAddress string) Handler {
	return &PublicPostHandler{
		postClientAddress: postClientAddress,
		userClientAddress: userClientAddress,
	}
}

func (handler *PublicPostHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/user/{id}/public", handler.GetPublicPosts)
	if err != nil {
		fmt.Println("Panikaaa")
		panic(err)
	}
}

func (handler *PublicPostHandler) GetPublicPosts(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userStatus := &domain.UserStatusRequest{Id: id}

	err := handler.checkUserStatus(userStatus)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if userStatus.IsPrivate == true {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		err1 := handler.getPosts(userStatus)
		if err1 != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		response, err := json.Marshal(userStatus.Posts)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)

	}
	return
}

func (handler *PublicPostHandler) checkUserStatus(userPostsRequest *domain.UserStatusRequest) error {
	userClient := services.NewUserClient(handler.userClientAddress)
	userProfileStatus, err := userClient.IsPrivate(context.TODO(), &user.IsPrivateRequest{Id: userPostsRequest.Id})
	if err != nil {
		return err
	}
	userPostsRequest.IsPrivate = userProfileStatus.Private
	return nil
}

func (handler *PublicPostHandler) getPosts(userPostsRequest *domain.UserStatusRequest) error {
	postClient := services.NewPostClient(handler.postClientAddress)
	postCollection, err := postClient.GetAllFromCollection(context.TODO(), &post.GetRequest{Id: userPostsRequest.Id})
	if err != nil {
		return err
	}

	userPostsRequest.Posts = postCollection.Posts
	return nil
}
