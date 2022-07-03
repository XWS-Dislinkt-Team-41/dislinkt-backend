package api

import (
	pb "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/job_offer_service"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/job_offer_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapJobOffer(jobOffer *domain.JobOffer) *pb.JobOffer {
	jobOfferPb := &pb.JobOffer{
		Id:            jobOffer.Id.Hex(),
		UserId:        jobOffer.UserId.Hex(),
		Position:      jobOffer.Position,
		Seniority:	   jobOffer.Seniority,
		Description:   jobOffer.Description,
		Prerequisites: jobOffer.Prerequisites,
	}
	return jobOfferPb
}

func mapNewJobOffer(jobOfferPb *pb.JobOffer) *domain.JobOffer {
	id, _ := primitive.ObjectIDFromHex(jobOfferPb.Id)
	userId, _ := primitive.ObjectIDFromHex(jobOfferPb.UserId)

	jobOffer := &domain.JobOffer{
		Id:            id,
		UserId:        userId,
		Position:      jobOfferPb.Position,
		Seniority:	   jobOfferPb.Seniority,
		Description:   jobOfferPb.Description,
		Prerequisites: jobOfferPb.Prerequisites,
	}
	return jobOffer
}
