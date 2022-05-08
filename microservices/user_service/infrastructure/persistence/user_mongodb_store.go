package persistence

import (
	"context"
	"errors"

	"strings"

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
	filter := bson.M{"_id": id}
	result, error := store.filterOne(filter)
	return result.IsPrivate, error
}

func (store *UserMongoDBStore) GetAll() ([]*domain.User, error) {
	filter := bson.D{}
	return store.filter(filter)
}

func (store *UserMongoDBStore) GetAllPublicUserId() ([]primitive.ObjectID, error) {
	var filteredUsers []*domain.User
	users, err := store.users.Find(context.TODO(), bson.M{"isPrivate": false})
	if err != nil {
		return nil, err
	}
	var ids []primitive.ObjectID
	if err = users.All(context.TODO(), &filteredUsers); err != nil {
		return nil, err
	}
	for _, user := range filteredUsers {
		ids = append(ids, user.Id)
	}
	return ids, err
}

func (store *UserMongoDBStore) SearchPublic(filter string) ([]*domain.User, error) {
	var foundUsers []*domain.User

	filter = strings.TrimSpace(filter)
	splitSearch := strings.Split(filter, " ")

	for _, splitSearchpart := range splitSearch {

		//username
		filtereds, err := store.users.Find(context.TODO(), bson.M{"isPrivate": false, "username": primitive.Regex{Pattern: splitSearchpart, Options: "i"}})
		if err != nil {
			return nil, err
		}
		var usersUsername []*domain.User
		if err = filtereds.All(context.TODO(), &usersUsername); err != nil {
			return nil, err
		}
		for _, userOneSlice := range usersUsername {
			foundUsers = AppendIfMissing(foundUsers, userOneSlice)
		}

		//name
		filtereds, err = store.users.Find(context.TODO(), bson.M{"isPrivate": false, "firstname": primitive.Regex{Pattern: splitSearchpart, Options: "i"}})
		if err != nil {
			return nil, err
		}
		var usersName []*domain.User
		if err = filtereds.All(context.TODO(), &usersName); err != nil {
			return nil, err
		}
		for _, userOneSlice := range usersName {
			foundUsers = AppendIfMissing(foundUsers, userOneSlice)
		}

		//surname
		filtereds, err = store.users.Find(context.TODO(), bson.M{"isPrivate": false, "lastname": primitive.Regex{Pattern: splitSearchpart, Options: "i"}})
		if err != nil {
			return nil, err
		}
		var usersSurname []*domain.User
		if err = filtereds.All(context.TODO(), &usersSurname); err != nil {
			return nil, err
		}
		for _, userOneSlice := range usersSurname {
			foundUsers = AppendIfMissing(foundUsers, userOneSlice)
		}
	}
	return foundUsers, nil
}

func AppendIfMissing(slice []*domain.User, i *domain.User) []*domain.User {
	for _, ele := range slice {
		if ele.Id == i.Id {
			return slice
		}
	}
	return append(slice, i)
}

func (store *UserMongoDBStore) Insert(user *domain.User) (*domain.User, error) {
	filter := bson.M{"username": user.Username}
	userInDatabase, _ := store.filterOneRegister(filter)

	user.Id = primitive.NewObjectID()
	if userInDatabase != nil {
		return nil, errors.New("user with the same id already exists")
	}
	// userInDatabase, _ = store.GetByEmail(user.Email)
	// if userInDatabase != nil {
	// 	return nil, errors.New("user with this email has already been registered")
	// }
	// userInDatabase, _ = store.GetByUsername(user.Username)
	// if userInDatabase != nil {
	// 	return nil, errors.New("username is taken")
	// }
	_, err1 := store.users.InsertOne(context.TODO(), user)
	if err1 != nil {
		return nil, errors.New("register error")
	}

	return user, nil
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

func (store *UserMongoDBStore) filterOneRegister(filter interface{}) (user *domain.User, err error) {
	result := store.users.FindOne(context.TODO(), filter)
	err = result.Decode(&user)
	if err != nil {
		return nil, nil
	}
	return
}
