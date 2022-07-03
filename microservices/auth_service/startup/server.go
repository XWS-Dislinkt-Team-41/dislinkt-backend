package startup

import (
	"fmt"
	"log"
	"net"

	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/auth_service/application"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/auth_service/domain"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/auth_service/infrastructure/api"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/auth_service/infrastructure/persistence"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/auth_service/startup/config"
	auth_service "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/auth_service"
	saga "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/saga/messaging"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/saga/messaging/nats"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "auth_service"
)

func (server *Server) Start() {

	mongoClient := server.initMongoClient()
	authStore := server.initAuthStore(mongoClient)

	commandPublisher := server.initPublisher(server.config.RegisterUserCommandSubject)
	replySubscriber := server.initSubscriber(server.config.RegisterUserReplySubject, QueueGroup)
	createOrderOrchestrator := server.initRegisterUserOrchestrator(commandPublisher, replySubscriber)

	authService := server.initAuthService(authStore, createOrderOrchestrator)

	commandSubscriber := server.initSubscriber(server.config.RegisterUserCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.RegisterUserReplySubject)
	server.initRegisterUserHandler(authService, replyPublisher, commandSubscriber)

	authHandler := server.initAuthHandler(authService)
	server.startGrpcServer(authHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.AuthDBHost, server.config.AuthDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initAuthStore(client *mongo.Client) domain.AuthStore {
	store := persistence.NewAuthMongoDBStore(client)
	store.DeleteAll()
	return store
}

func (server *Server) initRegisterUserOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *application.RegisterUserOrchestrator {
	orchestrator, err := application.NewRegisterUserOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) initAuthService(store domain.AuthStore, orchestrator *application.RegisterUserOrchestrator) *application.AuthService {
	return application.NewAuthService(store, orchestrator)
}

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initRegisterUserHandler(service *application.AuthService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewRegisterUserCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initAuthHandler(service *application.AuthService) *api.AuthHandler {
	return api.NewAuthHandler(service)
}

func (server *Server) startGrpcServer(authHandler *api.AuthHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	auth_service.RegisterAuthServiceServer(grpcServer, authHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
