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
	jobOfferGw "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/job_offer_service"
	messageGw "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/message_service"
	notificationGw "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/notification_service"
	postGw "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/post_service"
	userGw "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/rs/cors"
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
	jobOfferEndpoint := fmt.Sprintf("%s:%s", server.config.JobOfferHost, server.config.JobOfferPort)
	err = jobOfferGw.RegisterJobOfferServiceHandlerFromEndpoint(context.TODO(), server.mux, jobOfferEndpoint, opts)
	if err != nil {
		panic(err)
	}
	messageEndpoint := fmt.Sprintf("%s:%s", server.config.MessageHost, server.config.MessagePort)
	err = messageGw.RegisterMessageServiceHandlerFromEndpoint(context.TODO(), server.mux, messageEndpoint, opts)
	notificationEndpoint := fmt.Sprintf("%s:%s", server.config.NotificationHost, server.config.NotificationPort)
	err = notificationGw.RegisterNotificationServiceHandlerFromEndpoint(context.TODO(), server.mux, notificationEndpoint, opts)
	if err != nil {
		panic(err)
	}
}

func (server *Server) initCustomHandlers() {
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	postEndpoint := fmt.Sprintf("%s:%s", server.config.PostHost, server.config.PostPort)
	connEndpoint := fmt.Sprintf("%s:%s", server.config.ConnectHost, server.config.ConnectPort)
	publicPostHandler := api.NewPublicPostHandler(userEndpoint, postEndpoint)
	publicPostHandler.Init(server.mux)
	connectPostHandler := api.NewConnectPostHandler(postEndpoint, connEndpoint)
	connectPostHandler.Init(server.mux)
}

func (server *Server) Start() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(middleware.IsAuthenticated(server.mux))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), handler))
}
