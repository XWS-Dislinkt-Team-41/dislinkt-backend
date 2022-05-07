package startup

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/api_gateway/infrastructure/api"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/api_gateway/infrastructure/middleware"
	cfg "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/api_gateway/startup/config"
	authGw "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/auth_service"
	connectGw "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/connect_service"
	postGw "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/post_service"
	userGw "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	server.initCustomHandlers()
	return server
}

func (server *Server) initHandlers() {

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	postEmdpoint := fmt.Sprintf("%s:%s", server.config.PostHost, server.config.PostPort)
	err := postGw.RegisterPostServiceHandlerFromEndpoint(context.TODO(), server.mux, postEmdpoint, opts)
	if err != nil {
		panic(err)
	}
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	err = userGw.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)
	if err != nil {
		panic(err)
	}
	authEmdpoint := fmt.Sprintf("%s:%s", server.config.AuthHost, server.config.AuthPort)
	err = authGw.RegisterAuthServiceHandlerFromEndpoint(context.TODO(), server.mux, authEmdpoint, opts)
	if err != nil {

		panic(err)
	}
	connectEmdpoint := fmt.Sprintf("%s:%s", server.config.ConnectHost, server.config.ConnectPort)
	err = connectGw.RegisterConnectServiceHandlerFromEndpoint(context.TODO(), server.mux, connectEmdpoint, opts)
	if err != nil {
		panic(err)
	}
}

func (server *Server) initCustomHandlers() {
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	postEndpoint := fmt.Sprintf("%s:%s", server.config.PostHost, server.config.PostPort)
	authEndpoint := fmt.Sprintf("%s:%s", server.config.AuthHost, server.config.AuthPort)
	publicPostHandler := api.NewPublicPostHandler(userEndpoint, postEndpoint)
	publicPostHandler.Init(server.mux)
	registerHandler := api.NewRegisterHandler(userEndpoint, authEndpoint)
	registerHandler.Init(server.mux)
}

func (server *Server) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), middleware.IsAuthenticated(server.mux)))
}
