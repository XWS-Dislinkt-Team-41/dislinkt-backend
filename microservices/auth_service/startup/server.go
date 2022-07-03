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

func (server *Server) Start() {

	mongoClient := server.initMongoClient()
	authStore := server.initAuthStore(mongoClient)
	permissionStore := server.initPermissionStore(mongoClient)
	authService := server.initAuthService(authStore,permissionStore)
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
	for _, credential := range credentials {
		_, err := store.Insert(credential)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initAuthService(authStore domain.AuthStore, permissionStore domain.PermissionStore) *application.AuthService {
	return application.NewAuthService(authStore, permissionStore)
}

func (server *Server) initAuthHandler(service *application.AuthService) *api.AuthHandler {
	return api.NewAuthHandler(service)
}

func (server *Server) initPermissionStore(client *mongo.Client) domain.PermissionStore {
	store := persistence.NewPermissionMongoDBStore(client)
	store.DeleteAll()
	for _, permission := range permissions {
		_, err := store.Insert(permission)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
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
