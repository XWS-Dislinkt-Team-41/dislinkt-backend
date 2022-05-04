package persistence

import (
	"context"
	"log"

	"errors"
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
	x := []string{"000000000000000000000000", "110000000000000000000000"}
	posts := []*domain.Post{}
	for _, id := range x {
		userPost, _ := store.filter(filter, id)
		for _, post := range userPost {
			posts = append(posts, post)
		}
	}
	return posts, nil
}

func (store *PostMongoDBStore) GetAllFromCollection(id primitive.ObjectID) (post []*domain.Post, err error) {
	filter := bson.D{{}}
	return store.filter(filter, id.Hex())
}

func (store *PostMongoDBStore) Insert(id primitive.ObjectID, post *domain.Post) (*domain.Post, error) {

	insertResult, err := store.dbPost.Collection(COLLECTION+id.Hex()).InsertOne(context.TODO(), &domain.Post{
		Id:      primitive.NewObjectID(),
		Text:    post.Text,
		Link:    post.Link,
		Image:   post.Image,
		OwnerId: id,
	})
	if err != nil {
		log.Fatal(err)
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

func (store *PostMongoDBStore) UpdateLikes(reaction *domain.Reaction) (*domain.Post, error) {
	post, err := store.Get(reaction.Id, reaction.PostId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if contains(post.LikedBy, reaction.ReactionBy.Hex()) {
		return nil, errors.New("User already liked post")
	}
	post.LikedBy = append(post.LikedBy, reaction.ReactionBy.Hex())
	post.Likes = post.Likes + 1

	filter := bson.M{"_id": reaction.PostId}
	update := bson.D{
		{"$set", bson.D{{"liked_by", post.LikedBy}, {"likes", post.Likes}}},
	}

	insertResult, err := store.dbPost.Collection(COLLECTION+reaction.Id.Hex()).UpdateOne(context.TODO(), filter,
		update)
	if err != nil {
		return nil, err
	}
	if insertResult.MatchedCount != 1 {
		log.Fatal(err, "one document should've been updated")
		return nil, err
	}
	return post, err

}

func (store *PostMongoDBStore) UpdateDislikes(reaction *domain.Reaction) (*domain.Post, error) {
	post, err := store.Get(reaction.Id, reaction.PostId)
	if err != nil {
		log.Fatal(err)
	}
	if contains(post.DislikedBy, reaction.ReactionBy.Hex()) {
		return nil, errors.New("User already disliked post")
	}
	post.DislikedBy = append(post.DislikedBy, reaction.ReactionBy.Hex())
	post.Dislikes = post.Dislikes + 1

	filter := bson.M{"_id": reaction.PostId}
	update := bson.D{
		{"$set", bson.D{{"disliked_by", post.DislikedBy}, {"dislikes", post.Dislikes}}},
	}

	insertResult, err := store.dbPost.Collection(COLLECTION+reaction.Id.Hex()).UpdateOne(context.TODO(), filter,
		update)
	if err != nil {
		return nil, err
	}
	if insertResult.MatchedCount != 1 {
		log.Fatal(err, "one document should've been updated")
		return nil, err
	}
	return post, err

}

func (store *PostMongoDBStore) DeleteAll() {
	result, err := store.dbPost.ListCollectionNames(
		context.TODO(),
		bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	for _, coll := range result {
		store.dbPost.Collection(coll).DeleteMany(context.TODO(), bson.D{{}})
	}

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

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
