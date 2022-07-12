package persistence

import (
	"context"
	"errors"

	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/auth_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	PERMISSION_COLLECTION = "permissions"
)

type PermissionMongoDBStore struct {
	permissions *mongo.Collection
}

func NewPermissionMongoDBStore(client *mongo.Client) domain.PermissionStore {
	permissions := client.Database(DATABASE).Collection(PERMISSION_COLLECTION)
	return &PermissionMongoDBStore{
		permissions: permissions,
	}
}

func (store *PermissionMongoDBStore) Get(id primitive.ObjectID) (*domain.Permission, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *PermissionMongoDBStore) GetByRole(role domain.Role) ([]*domain.Permission, error) {
	filter := bson.M{"role": role}
	return store.filter(filter)
}

func (store *PermissionMongoDBStore) GetAll() ([]*domain.Permission, error) {
	filter := bson.D{}
	return store.filter(filter)
}


func (store *PermissionMongoDBStore) Insert(permission *domain.Permission) (*domain.Permission, error) {
	filter := bson.M{"id": permission.Id}
	permissionInDatabase, _ := store.filterOne(filter)
	if permissionInDatabase != nil {
		return nil, errors.New("Permission with the same id already exists.")
	}
	permission.Id = primitive.NewObjectID()
	_, err := store.permissions.InsertOne(context.TODO(), permission)
	if err != nil {
		return nil, errors.New("Create error.")
	}

	return permission, nil
}

func (store *PermissionMongoDBStore) Update(permission *domain.Permission) (*domain.Permission, error) {
	permissionInDatabase, err := store.Get(permission.Id)
	if permissionInDatabase == nil {
		return nil, err
	}
	permissionInDatabase.Role = permission.Role
	permissionInDatabase.Url = permission.Url
	filter := bson.M{"_id": permissionInDatabase.Id}
	update := bson.M{
		"$set": permissionInDatabase,
	}
	_, err = store.permissions.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return permissionInDatabase, nil
}

func (store *PermissionMongoDBStore) DeleteAll() {
	store.permissions.DeleteMany(context.TODO(), bson.D{})
}

func (store *PermissionMongoDBStore) filterOne(filter interface{}) (permission *domain.Permission, err error) {
	result := store.permissions.FindOne(context.TODO(), filter)
	err = result.Decode(&permission)
	if err != nil {
		return nil, nil
	}
	return
}

func (store *PermissionMongoDBStore) checkIfExsist(filter interface{}) (exsist bool, err error) {
	_, err = store.permissions.Find(context.TODO(), filter)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (store *PermissionMongoDBStore) filter(filter interface{}) ([]*domain.Permission, error) {
	cursor, err := store.permissions.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func decode(cursor *mongo.Cursor) (permissions []*domain.Permission, err error) {
	for cursor.Next(context.TODO()) {
		var Permission domain.Permission
		err = cursor.Decode(&Permission)
		if err != nil {
			return
		}
		permissions = append(permissions, &Permission)
	}
	err = cursor.Err()
	return
}
