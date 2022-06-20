package startup

import (
	"fmt"
	"log"
	"net"

	connections "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/connect_service"
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

func (server *Server) Start() {
	neo4jClient := server.initNeo4jClient()

	connectStore := server.initConnectStore(neo4jClient)

	connectService := server.initConnectService(connectStore)

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
	return store
}

func (server *Server) initConnectService(store domain.ConnectStore) *application.ConnectService {
	return application.NewConnectService(store)
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
