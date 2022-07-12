package startup

import (
	"fmt"
	"log"
	"net"

	job_offer "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/job_offer_service"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/job_offer_service/application"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/job_offer_service/domain"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/job_offer_service/infrastructure/api"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/job_offer_service/infrastructure/persistence"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/job_offer_service/startup/config"
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
	QueueGroup = "job_offer_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	jobOfferStore := server.initJobOfferStore(mongoClient)

	jobOfferService := server.initJobOfferService(jobOfferStore)

	jobOfferHandler := server.initJobOfferHandler(jobOfferService)

	server.startGrpcServer(jobOfferHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.JobOfferDBHost, server.config.JobOfferDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initJobOfferStore(client *mongo.Client) domain.JobOfferStore {
	store := persistence.NewJobOfferMongoDBStore(client)
	store.DeleteAll()
	for _, jobOffer := range jobOffers {
		_, err := store.Insert(jobOffer)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initJobOfferService(store domain.JobOfferStore) *application.JobOfferService {
	return application.NewJobOfferService(store)
}

func (server *Server) initJobOfferHandler(service *application.JobOfferService) *api.JobOfferHandler {
	return api.NewJobOfferHandler(service)
}

func (server *Server) startGrpcServer(jobOfferHandler *api.JobOfferHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	job_offer.RegisterJobOfferServiceServer(grpcServer, jobOfferHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
