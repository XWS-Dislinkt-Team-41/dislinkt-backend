package persistence

import (
	"context"

	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/post_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "post"
	COLLECTION = "posts"
)

type PostMongoDBStore struct {
	dbPost *mongo.Database
}

func NewPostMongoDBStore(client *mongo.Client) domain.PostStore {

	dbPost := client.Database(DATABASE)
	return &PostMongoDBStore{
		dbPost: dbPost,
	}
}

func (store *PostMongoDBStore) Get(id primitive.ObjectID) (post *domain.Post, err error) {

	filter := bson.M{"_id": id}
	posts := store.dbPost.Collection(COLLECTION + id.Hex())
	result := posts.FindOne(context.TODO(), filter)
	err = result.Decode(&post)
	return
}

func (store *PostMongoDBStore) GetAll() ([]*domain.Post, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *PostMongoDBStore) Insert(post *domain.Post) (*domain.Post, error) {

	posts := store.dbPost.Collection(COLLECTION + "623b0cc3a34d25d8567f9f82")
	result, err := posts.InsertOne(context.TODO(), post)
	if err != nil {
		return nil, err
	}
	post.Id = result.InsertedID.(primitive.ObjectID)
	return post, nil
}

func (store *PostMongoDBStore) DeleteAll() {
	// posts := store.dbPost.Collection(COLLECTION+"123")
	// posts.DeleteMany(context.TODO(), bson.D{{}})

	store.dbPost.Collection(COLLECTION+"623b0cc3a34d25d8567f9f82").DeleteMany(context.TODO(), bson.D{{}})
}

func (store *PostMongoDBStore) filter(filter interface{}) ([]*domain.Post, error) {
	posts := store.dbPost.Collection(COLLECTION + "623b0cc3a34d25d8567f9f82")
	cursor, err := posts.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func decode(cursor *mongo.Cursor) (posts []*domain.Post, err error) {
	for cursor.Next(context.TODO()) {
		var post domain.Post
		err = cursor.Decode(&post)
		if err != nil {
			return
		}
		posts = append(posts, &post)
	}
	err = cursor.Err()
	return
}
