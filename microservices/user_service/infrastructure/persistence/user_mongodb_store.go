package persistence

import (
	"context"
	"fmt"

	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "user"
	COLLECTION = "user"
)

type UserMongoDBStore struct {
	users *mongo.Collection
}

func NewUserMongoDBStore(client *mongo.Client) domain.UserStore {
	users := client.Database(DATABASE).Collection(COLLECTION)
	return &UserMongoDBStore{
		users: users,
	}
}

func (store *UserMongoDBStore) GetByEmail(email string) (*domain.User, error) {
	filter := bson.M{"email": email}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) GetByUsername(username string) (*domain.User, error) {
	filter := bson.M{"username": username}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) Get(id primitive.ObjectID) (*domain.User, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) IsPrivate(id primitive.ObjectID) (bool, error) {
	filter := bson.M{"_id":id}
	result, error := store.filterOne(filter)
	return result.IsPrivate, error
}

func (store *UserMongoDBStore) GetAll() ([]*domain.User, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *UserMongoDBStore) SearchPublic(username string, name string) ([]*domain.User, error) {
	filter := bson.M{"isPrivate": false, "$or": []bson.M{bson.M{"username": username}, bson.M{"firstname": name}}}
	return store.filter(filter)
}

func (store *UserMongoDBStore) Insert(user *domain.User) (string, error) {
	userInDatabase, err := store.Get(user.Id)
	user.Id = primitive.NewObjectID()
	if userInDatabase != nil {
		return "User with the same id already exists.", nil
	}
	userInDatabase, err = store.GetByEmail(user.Email)
	if userInDatabase != nil {
		return "User with this email has already been registered.", nil
	}
	userInDatabase, err = store.GetByUsername(user.Username)
	if userInDatabase != nil {
		return "Username is taken.", nil
	}
	result, err := store.users.InsertOne(context.TODO(), user)
	if err != nil {
		return "Register error.", err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)
	return "User has been registered.", nil
}

func (store *UserMongoDBStore) DeleteAll() {
	store.users.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *UserMongoDBStore) filter(filter interface{}) ([]*domain.User, error) {
	cursor, err := store.users.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *UserMongoDBStore) UpdatePersonalInfo(user *domain.User) (string, error) {
	fmt.Println(user.Id)
	userInDatabase, err := store.Get(user.Id)
	if userInDatabase == nil {
		return "User doesn't exist.", nil
	}
	checkUsername, err := store.GetByUsername(user.Username)
	if checkUsername != nil {
		if checkUsername.Id != userInDatabase.Id {
			return "Username is taken.", nil
		}
	}
	userInDatabase.Firstname = user.Firstname
	userInDatabase.Email = user.Email
	userInDatabase.MobileNumber = user.MobileNumber
	userInDatabase.Gender = user.Gender
	userInDatabase.BirthDay = user.BirthDay
	userInDatabase.Username = user.Username
	userInDatabase.Biography = user.Biography
	filter := bson.M{"_id": userInDatabase.Id}
	update := bson.M{
		"$set": userInDatabase,
	}
	_, err = store.users.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "Update failed.", err
	}

	return "User has been updated.", nil
}

func (store *UserMongoDBStore) UpdateCareerInfo(user *domain.User) (string, error) {
	userInDatabase, err := store.Get(user.Id)
	if userInDatabase == nil {
		return "User doesn't exist.", nil
	}
	userInDatabase.Experience = user.Experience
	userInDatabase.Education = user.Education
	filter := bson.M{"_id": userInDatabase.Id}
	update := bson.M{
		"$set": userInDatabase,
	}
	_, err = store.users.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "Update failed.", err
	}

	return "User has been updated.", nil
}

func (store *UserMongoDBStore) UpdateInterestsInfo(user *domain.User) (string, error) {

	userInDatabase, err := store.Get(user.Id)
	if userInDatabase == nil {
		return "user doesn't exist", nil
	}
	userInDatabase.Skills = user.Skills
	userInDatabase.Interests = user.Interests
	filter := bson.M{"_id": userInDatabase.Id}
	update := bson.M{
		"$set": userInDatabase,
	}
	_, err = store.users.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return "Update failed.", err
	}

	return "User has been updated.", nil
}

func (store *UserMongoDBStore) filterOne(filter interface{}) (User *domain.User, err error) {
	result := store.users.FindOne(context.TODO(), filter)
	err = result.Decode(&User)
	return
}

func decode(cursor *mongo.Cursor) (users []*domain.User, err error) {
	for cursor.Next(context.TODO()) {
		var User domain.User
		err = cursor.Decode(&User)
		if err != nil {
			return
		}
		users = append(users, &User)
	}
	err = cursor.Err()
	return
}
