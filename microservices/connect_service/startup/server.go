package startup

import (
	"fmt"
	"log"
	"net"

	connections "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/connect_service"
	saga "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/saga/messaging"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/saga/messaging/nats"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/application"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/domain"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/infrastructure/api"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/infrastructure/persistence"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/startup/config"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
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
	QueueGroup = "connect_service"
)

func (server *Server) Start() {
	neo4jClient := server.initNeo4jClient()

	connectStore := server.initConnectStore(neo4jClient)

	connectService := server.initConnectService(connectStore)

	commandSubscriber := server.initSubscriber(server.config.RegisterUserCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.RegisterUserReplySubject)
	server.initRegisterUserHandler(connectService, replyPublisher, commandSubscriber)

	commandSubscriber = server.initSubscriber(server.config.ChangeAccountPrivacyCommandSubject, QueueGroup)
	replyPublisher = server.initPublisher(server.config.ChangeAccountPrivacyReplySubject)
	server.initChangePrivacyHandler(connectService, replyPublisher, commandSubscriber)
	
	connectHandler := server.initConnectHandler(connectService)

	server.startGrpcServer(connectHandler)
}

func (server *Server) initNeo4jClient() *neo4j.Driver {
	driver, err := persistence.GetDriver(server.config.ConnectDBHost, server.config.ConnectDBUser, server.config.ConnectDBPass, server.config.ConnectDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return driver
}

func (server *Server) initConnectStore(driver *neo4j.Driver) domain.ConnectStore {
	store := persistence.NewConnectNeo4jDBStore(driver)
	err := store.InitNeo4jDB()
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		_, err := store.Register(*user)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initConnectService(store domain.ConnectStore) *application.ConnectService {
	return application.NewConnectService(store)
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

func (server *Server) initRegisterUserHandler(service *application.ConnectService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewRegisterUserCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initChangePrivacyHandler(service *application.ConnectService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewChangePrivacyCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initConnectHandler(service *application.ConnectService) *api.ConnectHandler {
	return api.NewConnectHandler(service)
}

func (server *Server) startGrpcServer(connectHandler *api.ConnectHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	connections.RegisterConnectServiceServer(grpcServer, connectHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
