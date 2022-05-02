package persistence

import (
	"context"
	"log"

	"fmt"

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

func (store *PostMongoDBStore) Get(id primitive.ObjectID, post_id primitive.ObjectID) (post *domain.Post, err error) {

	filter := bson.M{"_id": post_id}
	posts := store.dbPost.Collection(COLLECTION + id.Hex())
	result := posts.FindOne(context.TODO(), filter)
	err = result.Decode(&post)
	return
}

func (store *PostMongoDBStore) GetAll() ([]*domain.Post, error) {
	filter := bson.D{{}}
	return store.filter(filter, "000000000000000000000000")
}

func (store *PostMongoDBStore) GetAllFromCollection(id primitive.ObjectID) (post []*domain.Post, err error) {

	filter := bson.D{{}}
	return store.filter(filter, id.Hex())
}

func (store *PostMongoDBStore) Insert(id primitive.ObjectID, post *domain.Post) (*domain.Post, error) {

	insertResult, err := store.dbPost.Collection(COLLECTION+id.Hex()).InsertOne(context.TODO(), &domain.Post{
		Id:       primitive.NewObjectID(),
		Text:     post.Text,
		Link:     post.Link,
		Image:    post.Image,
		OwnerId:  post.OwnerId,
		Likes:    post.Likes,
		Dislikes: post.Dislikes,
	})
	if err != nil {
		log.Fatal(err, "Greskaaaa")
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	return post, nil
}

func (store *PostMongoDBStore) InsertComment(id primitive.ObjectID, post_id primitive.ObjectID, comment *domain.Comment) (*domain.Comment, error) {
	post, err := store.Get(id, post_id)
	if err != nil {
		log.Fatal(err)
	}
	post.Comments = append(post.Comments, *comment)

	filter := bson.M{"_id": post_id}
	update := bson.D{
		{"$set", bson.D{{"comments", post.Comments}}},
	}

	insertResult, err := store.dbPost.Collection(COLLECTION+id.Hex()).UpdateOne(context.TODO(), filter,
		update)
	if err != nil {
		return nil, err
	}
	if insertResult.MatchedCount != 1 {
		log.Fatal(err, "one document should've been updated")
		return nil, err
	}
	return comment, err

}

func (store *PostMongoDBStore) DeleteAll() {

	store.dbPost.Collection(COLLECTION+"000000000000000000000000").DeleteMany(context.TODO(), bson.D{{}})
}

func (store *PostMongoDBStore) filter(filter interface{}, id string) ([]*domain.Post, error) {
	posts := store.dbPost.Collection(COLLECTION + id)
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
