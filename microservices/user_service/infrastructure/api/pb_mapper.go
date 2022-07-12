package api

import (
	"time"

	pb "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/user_service"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service/domain"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service/domain/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
	events "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/saga/change_account_privacy"
)

func mapUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:           user.Id.Hex(),
		Firstname:    user.Firstname,
		Lastname:     user.Lastname,
		Email:        user.Email,
		MobileNumber: user.MobileNumber,
		Gender:       mapGender(user.Gender),
		BirthDay:     timestamppb.New(user.BirthDay),
		Username:     user.Username,
		Biography:    user.Biography,
		Experience:   user.Experience,
		Education:    mapEducation(user.Education),
		Skills:       user.Skills,
		Interests:    user.Interests,
		Password:     user.Password,
		IsPrivate:    user.IsPrivate,
	}
	return userPb
}

func mapNewUser(userPb *pb.User) *domain.User {
	id, _ := primitive.ObjectIDFromHex(userPb.Id)
	if userPb.BirthDay != nil {
		user := &domain.User{
			Id:           id,
			Firstname:    userPb.Firstname,
			Lastname:     userPb.Lastname,
			Email:        userPb.Email,
			MobileNumber: userPb.MobileNumber,
			Gender:       mapNewGender(userPb.Gender),
			BirthDay:     userPb.BirthDay.AsTime(),
			Username:     userPb.Username,
			Biography:    userPb.Biography,
			Experience:   userPb.Experience,
			Education:    mapNewEducation(userPb.Education),
			Skills:       userPb.Skills,
			Interests:    userPb.Interests,
			Password:     userPb.Password,
			IsPrivate:    userPb.IsPrivate,
		}
		return user
	} else {
		user := &domain.User{
			Id:           id,
			Firstname:    userPb.Firstname,
			Lastname:     userPb.Lastname,
			Email:        userPb.Email,
			MobileNumber: userPb.MobileNumber,
			Gender:       mapNewGender(userPb.Gender),
			BirthDay:     time.Now(),
			Username:     userPb.Username,
			Biography:    userPb.Biography,
			Experience:   userPb.Experience,
			Education:    mapNewEducation(userPb.Education),
			Skills:       userPb.Skills,
			Interests:    userPb.Interests,
			Password:     userPb.Password,
			IsPrivate:    userPb.IsPrivate,
		}
		return user
	}
}

func mapPersonalInfoUser(userPb *pb.User) *domain.User {

	id, _ := primitive.ObjectIDFromHex(userPb.Id)

	if userPb.BirthDay != nil {
		user := &domain.User{
			Id:           id,
			Firstname:    userPb.Firstname,
			Lastname:     userPb.Lastname,
			Email:        userPb.Email,
			MobileNumber: userPb.MobileNumber,
			Gender:       mapNewGender(userPb.Gender),
			BirthDay:     userPb.BirthDay.AsTime(),
			Username:     userPb.Username,
			Biography:    userPb.Biography,
			Password:     userPb.Password,
		}
		return user
	} else {
		user := &domain.User{
			Id:           id,
			Firstname:    userPb.Firstname,
			Lastname:     userPb.Lastname,
			Email:        userPb.Email,
			MobileNumber: userPb.MobileNumber,
			Gender:       mapNewGender(userPb.Gender),
			BirthDay:     time.Now(),
			Username:     userPb.Username,
			Biography:    userPb.Biography,
			Password:     userPb.Password,
		}
		return user
	}
}

func mapCareerInfoUser(userPb *pb.User) *domain.User {
	id, _ := primitive.ObjectIDFromHex(userPb.Id)

	user := &domain.User{
		Id:         id,
		Experience: userPb.Experience,
		Education:  mapNewEducation(userPb.Education),
		Password:   userPb.Password,
	}
	return user
}

func mapInterestsInfoUser(userPb *pb.User) *domain.User {
	id, _ := primitive.ObjectIDFromHex(userPb.Id)

	user := &domain.User{
		Id:        id,
		Skills:    userPb.Skills,
		Interests: userPb.Interests,
		Password:  userPb.Password,
	}
	return user
}

func mapGender(status domain.Gender) pb.User_Gender {
	switch status {
	case domain.Male:
		return pb.User_Male
	}
	return pb.User_Female
}

func mapNewGender(status pb.User_Gender) domain.Gender {
	switch status {
	case pb.User_Male:
		return domain.Male
	}
	return domain.Female
}

func mapEducation(status enums.Education) pb.User_Education {
	switch status {
	case enums.Primary:
		return pb.User_Primary
	case enums.LowerSecondary:
		return pb.User_LowerSecondary
	case enums.UpperSecondary:
		return pb.User_UpperSecondary
	case enums.PostSecondary:
		return pb.User_PostSecondary
	case enums.ShortCycleTetriary:
		return pb.User_ShortCycleTetriary
	case enums.Bachelor:
		return pb.User_Bachelor
	case enums.Master:
		return pb.User_Master
	}
	return pb.User_Doctorate
}

func mapNewEducation(status pb.User_Education) enums.Education {
	switch status {
	case pb.User_Primary:
		return enums.Primary
	case pb.User_LowerSecondary:
		return enums.LowerSecondary
	case pb.User_UpperSecondary:
		return enums.UpperSecondary
	case pb.User_PostSecondary:
		return enums.PostSecondary
	case pb.User_ShortCycleTetriary:
		return enums.ShortCycleTetriary
	case pb.User_Bachelor:
		return enums.Bachelor
	case pb.User_Master:
		return enums.Master
	}
	return enums.Doctorate
}

func mapRegisterUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:        user.Id.Hex(),
		Username:  user.Username,
		IsPrivate: user.IsPrivate,
	}
	return userPb
}

func mapUserDetails(user *events.UserDetails) *pb.User {
	userPb := &pb.User{
		Id:           user.Id,
		Username:     user.Username,
		Password:     user.Password,
		IsPrivate:    user.IsPrivate,
		Firstname:    user.Firstname,
		Lastname:     user.Lastname,
		Email:        user.Email,
		MobileNumber: user.MobileNumber,
	}
	return userPb
}

func mapPbUserDetails(userPb *pb.User) *events.UserDetails {
	user := &events.UserDetails{
		Id:           userPb.Id,
		Username:     userPb.Username,
		Password:     userPb.Password,
		IsPrivate:    userPb.IsPrivate,
		Firstname:    userPb.Firstname,
		Lastname:     userPb.Lastname,
		Email:        userPb.Email,
		MobileNumber: userPb.MobileNumber,
	}
	return user
}

