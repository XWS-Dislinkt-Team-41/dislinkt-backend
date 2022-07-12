package api

import (
	"context"

	pb "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/job_offer_service"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/job_offer_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobOfferHandler struct {
	pb.UnimplementedJobOfferServiceServer
	service *application.JobOfferService
}

func NewJobOfferHandler(service *application.JobOfferService) *JobOfferHandler {
	return &JobOfferHandler{
		service: service,
	}
}

func (handler *JobOfferHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	jobOffer, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	jobOfferPb := mapJobOffer(jobOffer)
	response := &pb.GetResponse{
		JobOffer: jobOfferPb,
	}
	return response, nil
}

func (handler *JobOfferHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	jobOffers, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		JobOffers: []*pb.JobOffer{},
	}
	for _, jobOffer := range jobOffers {
		current := mapJobOffer(jobOffer)
		response.JobOffers = append(response.JobOffers, current)
	}
	return response, nil
}

func (handler *JobOfferHandler) Search(ctx context.Context, request *pb.SearchRequest) (*pb.SearchResponse, error) {
	filter := request.Filter
	jobOffers, err := handler.service.Search(filter)
	if err != nil {
		return nil, err
	}
	response := &pb.SearchResponse{
		JobOffers: []*pb.JobOffer{},
	}
	for _, jobOffer := range jobOffers {
		current := mapJobOffer(jobOffer)
		response.JobOffers = append(response.JobOffers, current)
	}
	return response, nil
}

func (handler *JobOfferHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	jobOffer := mapNewJobOffer(request.JobOffer)
	jobOffer, err := handler.service.Update(jobOffer)
	if err != nil {
		return nil, err
	}
	jobOfferPb := mapJobOffer(jobOffer)
	response := &pb.UpdateResponse{
		JobOffer: jobOfferPb,
	}
	return response, err
}

func (handler *JobOfferHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	jobOffer := mapNewJobOffer(request.JobOffer)
	jobOffer, err := handler.service.Insert(jobOffer)
	if err != nil {
		return nil, err
	}
	jobOfferPb := mapJobOffer(jobOffer)
	response := &pb.CreateResponse{
		JobOffer: jobOfferPb,
	}
	return response, err
}
