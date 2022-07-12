package persistence

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/message_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "chatroom"
	COLLECTION = "chatroom"
)

type MessageMongoDBStore struct {
	dbChatRoom *mongo.Database
}

func NewMessageMongoDBStore(client *mongo.Client) domain.MessageStore {
	dbChatRoom := client.Database(DATABASE)
	return &MessageMongoDBStore{
		dbChatRoom: dbChatRoom,
	}
}

func (store *MessageMongoDBStore) Get(id, connectedId primitive.ObjectID) ([]*domain.Message, error) {
	filter := bson.D{{}}
	chatRoomId := id.Hex() + "-" + connectedId.Hex()
	return store.filter(filter, chatRoomId)
}

func (store *MessageMongoDBStore) SendMessage(id, connectedId primitive.ObjectID, message *domain.Message) (*domain.Message, error) {

	insertResult, err := store.dbChatRoom.Collection(COLLECTION+id.Hex()+"-"+connectedId.Hex()).InsertOne(context.TODO(), &domain.Message{
		Id:       primitive.NewObjectID(),
		Text:     message.Text,
		SentTime: time.Now(),
		Seen:     message.Seen,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	return message, nil
}

func (store *MessageMongoDBStore) Insert(id, connectedId primitive.ObjectID, message *domain.Message) (*domain.Message, error) {

	insertResult, err := store.dbChatRoom.Collection(COLLECTION+id.Hex()+"-"+connectedId.Hex()).InsertOne(context.TODO(), &domain.Message{
		Id:       primitive.NewObjectID(),
		Text:     message.Text,
		SentTime: time.Now(),
		Seen:     message.Seen,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	return message, nil
}

func (store *MessageMongoDBStore) DeleteAll() {
	result, err := store.dbChatRoom.ListCollectionNames(
		context.TODO(),
		bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	for _, coll := range result {
		store.dbChatRoom.Collection(coll).DeleteMany(context.TODO(), bson.D{{}})
	}

}

func (store *MessageMongoDBStore) filter(filter interface{}, id string) ([]*domain.Message, error) {
	posts := store.dbChatRoom.Collection(COLLECTION + id)
	cursor, err := posts.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func decode(cursor *mongo.Cursor) (messages []*domain.Message, err error) {
	for cursor.Next(context.TODO()) {
		var Message domain.Message
		err = cursor.Decode(&Message)
		if err != nil {
			return
		}
		messages = append(messages, &Message)
	}
	err = cursor.Err()
	return
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
