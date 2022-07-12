package api

import (
	pb "github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/message_service"
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/message_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapChatRoom(chatRoom *domain.ChatRoom) *pb.ChatRoom {
	chatRoomPb := &pb.ChatRoom{
		Id:              chatRoom.Id.Hex(),
		Name:            chatRoom.Name,
		ParticipantsIds: chatRoom.ParticipantsIds,
	}
	for _, message := range chatRoom.Messages {
		chatRoomPb.Messages = append(chatRoomPb.Messages, &pb.Message{
			Id:       message.Id.Hex(),
			Text:     message.Text,
			SentTime: timestamppb.New(message.SentTime),
			Seen:     message.Seen,
		})
	}

	return chatRoomPb
}

func mapNewChatRoom(chatRoomPb *pb.ChatRoom) *domain.ChatRoom {
	id, _ := primitive.ObjectIDFromHex(chatRoomPb.Id)
	chatRoom := &domain.ChatRoom{
		Id:              id,
		Name:            chatRoomPb.Name,
		ParticipantsIds: chatRoomPb.ParticipantsIds,
	}
	for _, message := range chatRoomPb.Messages {
		messageId, _ := primitive.ObjectIDFromHex(message.Id)
		chatRoom.Messages = append(chatRoom.Messages, domain.Message{
			Id:       messageId,
			Text:     message.Text,
			SentTime: message.SentTime.AsTime(),
			Seen:     message.Seen,
		})
	}
	return chatRoom
}

func mapMessage(message *domain.Message) *pb.Message {
	messagePb := &pb.Message{
		Id:       message.Id.Hex(),
		Text:     message.Text,
		SentTime: timestamppb.New(message.SentTime),
		Seen:     message.Seen,
	}
	return messagePb
}

// func mapMessages(messages *[]domain.Message) *[]pb.Message {

// 	for _, message := range messages {
// 		messagesPb = append(messagesPb, &pb.Message{
// 			Id:       message.Id.Hex(),
// 			Text:     message.Text,
// 			SentTime: timestamppb.New(message.SentTime),
// 			Seen:     message.Seen,
// 		})
// 	}
// 	return messagesPb
// }
