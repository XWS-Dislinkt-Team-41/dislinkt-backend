package startup

import (
	"log"

	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/tree/develop/microservices/connect_service/infrastructure/persistence"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/tree/develop/microservices/connect_service/startup/config"
	"github.com/neo4j/neo4j-go-driver/neo4j"
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
	// neo4jClient := server.initNeo4jClient()
	// connectionStore := server.initconnectionStore(mongoClient)

	// connectionService := server.initconnectionService(connectionStore)

	// connectionHandler := server.initconnectionHandler(connectionService)

	// server.startGrpcServer(connectionHandler)
}

func (server *Server) initNeo4jClient() neo4j.Driver {
	client, err := persistence.GetClient(server.config.ConnectDBHost, server.config.ConnectDBUser, server.config.ConnectDBPass, server.config.ConnectDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// func (server *Server) initconnectionStore(client *mongo.Client) domain.connectionStore {
// 	store := persistence.NewconnectionMongoDBStore(client)
// 	store.DeleteAll()
// 	for _, connection := range connections {
// 		err := store.Insert(connection)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	return store
// }

// func (server *Server) initconnectionService(store domain.connectionStore) *application.connectionService {
// 	return application.NewconnectionService(store)
// }

// func (server *Server) initconnectionHandler(service *application.connectionService) *api.connectionHandler {
// 	return api.NewconnectionHandler(service)
// }

// func (server *Server) startGrpcServer(connectionHandler *api.connectionHandler) {
// 	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	grpcServer := grpc.NewServer()
// 	catalogue.RegisterCatalogueServiceServer(grpcServer, connectionHandler)
// 	if err := grpcServer.Serve(listener); err != nil {
// 		log.Fatalf("failed to serve: %s", err)
// 	}
// }
