package persistence

import (
	"fmt"

	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/domain"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	queryRegisterUser          = "CREATE (u:User{id:$userId, name:$userId, private:$userPrivate}) RETURN u"
	queryUpdateUser            = "MATCH (u:User{id:$userId}) SET u.private = $userPrivate RETURN u"
	queryIsUserPrivate         = "MATCH (u:User{id:$userId}) RETURN u.private AS IsUserPrivate"
	queryCreateConnection      = "MATCH (u1:User{id:$userId}) MATCH (u2:User{id:$cUserId}) MERGE (u1)-[c:Connection]-(u2)"
	queryDeleteConnection      = "MATCH (u1:User{id:$userId})-[c:Connection]-(u2:User{id:$cUserId}) DELETE c"
	queryGetConnections        = "MATCH (u:User{id:$userId}) MATCH (u)-[c:Connection]-(x) RETURN x.id"
	queryCreateInvite          = "MATCH (u1:User{id:$userId}) MATCH (u2:User{id:$cUserId}) MERGE (u1)-[i:Invite]->(u2)"
	queryDeleteInvite          = "MATCH (u1:User{id:$userId})-[i:Invite]->(u2:User{id:$cUserId}) DELETE i"
	queryDeleteReceivedInvite  = "MATCH (u1:User{id:$userId})<-[i:Invite]-(u2:User{id:$cUserId}) DELETE i"
	queryIsReceivedInvite      = "MATCH (u1:User{id:$userId}) MATCH (u2:User{id:$cUserId}) RETURN exists((u1)<-[:Invite]-(u2)) AS IsReceived"
	queryGetAllInvitations     = "MATCH (u:User{id:$userId}) MATCH (u)<-[i:Invite]-(x) RETURN x.id"
	queryGetAllSentInvitations = "MATCH (u:User{id:$userId}) MATCH (u)-[i:Invite]->(x) RETURN x.id"
)

type ConnectNeo4jDBStore struct {
	driver *neo4j.Driver
}

func NewConnectNeo4jDBStore(driver *neo4j.Driver) domain.ConnectStore {
	return &ConnectNeo4jDBStore{driver: driver}
}

func (store *ConnectNeo4jDBStore) Register(user domain.Profile) (*domain.Profile, error) {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		exists, err := store.IsUserExistTx(tx, user.Id)
		if err != nil {
			return nil, err
		}
		if !*exists {
			_, err = store.PersistUserTx(tx, user)
		} else {
			err = status.Error(codes.AlreadyExists, fmt.Sprint("user: %s already exist", user.Id.Hex()))
		}
		return nil, err
	})
	if err != nil {
		return nil, err
	}
	profile := domain.Profile{
		Id:      user.Id,
		Private: user.Private,
	}
	return &profile, nil
}

func (store *ConnectNeo4jDBStore) UpdateUser(user domain.Profile) (*domain.Profile, error) {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		exists, err := store.IsUserExistTx(tx, user.Id)
		if err != nil {
			return nil, err
		}
		if *exists {
			_, err = store.UpdateUserTx(tx, user)
		} else {
			err = status.Error(codes.InvalidArgument, fmt.Sprint("user: %s does not exist", user.Id.Hex()))
		}
		return nil, err
	})
	if err != nil {
		return nil, err
	}
	profile := domain.Profile{
		Id:      user.Id,
		Private: user.Private,
	}
	return &profile, nil
}

func (store *ConnectNeo4jDBStore) IsUserPrivate(userId primitive.ObjectID) (*bool, error) {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	var IsUserPrivate *bool
	_, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		exists, err := store.IsUserExistTx(tx, userId)
		if err != nil {
			return nil, err
		}
		if *exists {
			IsUserPrivate, err = store.IsUserPrivateTx(tx, userId)
		} else {
			err = status.Error(codes.InvalidArgument, fmt.Sprint("user: %s does not exist", userId.Hex()))
		}
		return nil, err
	})
	if err != nil {
		return nil, err
	}
	return IsUserPrivate, nil
}

func (store *ConnectNeo4jDBStore) Connect(userId, cUserId primitive.ObjectID) (*domain.Connection, error) {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := store.CreateConnectionTx(tx, userId, cUserId)
		return nil, err
	})
	if err != nil {
		return nil, err
	}
	connection := domain.Connection{
		User:  domain.Profile{Id: userId},
		CUser: domain.Profile{Id: cUserId},
	}
	return &connection, nil
}

func (store *ConnectNeo4jDBStore) UnConnect(userId, cUserId primitive.ObjectID) error {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		userExists, err := store.IsUserExistTx(tx, userId)
		if err != nil {
			return nil, err
		}
		cUserExists, err := store.IsUserExistTx(tx, cUserId)
		if err != nil {
			return nil, err
		}
		if *userExists && *cUserExists {
			_, err = store.DeleteConnectionTx(tx, userId, cUserId)
		} else {
			err = status.Error(codes.InvalidArgument, "user does not exist")
		}
		return nil, err
	})
	if err != nil {
		return err
	}
	return nil
}

func (store *ConnectNeo4jDBStore) GetUserConnections(userId primitive.ObjectID) ([]*domain.Connection, error) {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := store.GetConnectionsTx(tx, userId)
		return result, err
	})
	if err != nil {
		return nil, err
	}
	connections := result.([]*domain.Connection)
	return connections, nil
}

func (store *ConnectNeo4jDBStore) Invite(userId, cUserId primitive.ObjectID) (*domain.Connection, error) {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := store.CreateInviteTx(tx, userId, cUserId)
		return nil, err
	})
	if err != nil {
		return nil, err
	}
	invite := domain.Connection{
		User:  domain.Profile{Id: userId},
		CUser: domain.Profile{Id: cUserId},
	}
	return &invite, nil
}

func (store *ConnectNeo4jDBStore) AcceptInvitation(userId primitive.ObjectID, cUserId primitive.ObjectID) (*domain.Connection, error) {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	isCreatedConection := false
	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(queryIsReceivedInvite, map[string]interface{}{"userId": userId.Hex(), "cUserId": cUserId.Hex()})
		if err != nil {
			return nil, err
		}
		IsReceivedInvite := false
		for result.Next() {
			if value, ok := result.Record().Get("IsReceived"); ok {
				IsReceivedInvite = value.(bool)
			} else {
				return nil, err
			}
		}
		if IsReceivedInvite {
			_, err = transaction.Run(queryDeleteReceivedInvite, map[string]interface{}{"userId": userId.Hex(), "cUserId": cUserId.Hex()})
			if err != nil {
				return nil, err
			}
			result, err = transaction.Run(queryCreateConnection, map[string]interface{}{"userId": userId.Hex(), "cUserId": cUserId.Hex()})
			if err != nil {
				return nil, err
			}
			isCreatedConection = true
		}
		return result.Consume()
	})
	var connection *domain.Connection
	if isCreatedConection {
		connection = &domain.Connection{
			User:  domain.Profile{Id: userId},
			CUser: domain.Profile{Id: cUserId},
		}
	}
	if err != nil {
		return nil, err
	}
	return connection, nil
}

func (store *ConnectNeo4jDBStore) DeclineInvitation(userId primitive.ObjectID, cUserId primitive.ObjectID) error {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(queryDeleteReceivedInvite, map[string]interface{}{"userId": userId.Hex(), "cUserId": cUserId.Hex()})
		if err != nil {
			return nil, err
		}
		return nil, result.Err()
	})
	if err != nil {
		return err
	}
	return nil
}

func (store *ConnectNeo4jDBStore) CancelInvitation(userId primitive.ObjectID, cUserId primitive.ObjectID) error {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(queryDeleteInvite, map[string]interface{}{"userId": userId.Hex(), "cUserId": cUserId.Hex()})
		if err != nil {
			return nil, err
		}
		return nil, result.Err()
	})
	if err != nil {
		return err
	}
	return nil
}

func (store *ConnectNeo4jDBStore) GetAllInvitations(userId primitive.ObjectID) ([]*domain.Connection, error) {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	result, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(queryGetAllInvitations, map[string]interface{}{"userId": userId.Hex()})
		if err != nil {
			return nil, err
		}

		var invites []*domain.Connection
		var id string
		for result.Next() {
			// if value, ok := result.Record().GetByIndex(0).(string); ok {
			// 	id = value
			// } else {
			// 	return nil, err
			// }
			cUserId, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				return nil, err
			}
			invite := domain.Connection{
				User:  domain.Profile{Id: userId},
				CUser: domain.Profile{Id: cUserId},
			}
			invites = append(invites, &invite)
		}
		return invites, result.Err()
	})
	invites := result.([]*domain.Connection)
	if err != nil {
		return nil, err
	}
	return invites, nil
}

func (store *ConnectNeo4jDBStore) GetAllSentInvitations(userId primitive.ObjectID) ([]*domain.Connection, error) {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	result, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(queryGetAllSentInvitations, map[string]interface{}{"userId": userId.Hex()})
		if err != nil {
			return nil, err
		}

		var invites []*domain.Connection
		var id string
		for result.Next() {
			// if value, ok := result.Record().GetByIndex(0).(string); ok {
			// 	id = value
			// } else {
			// 	return nil, err
			// }
			cUserId, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				return nil, err
			}
			invite := domain.Connection{
				User:  domain.Profile{Id: userId},
				CUser: domain.Profile{Id: cUserId},
			}
			invites = append(invites, &invite)
		}
		return invites, result.Err()
	})
	invites := result.([]*domain.Connection)
	if err != nil {
		return nil, err
	}
	return invites, nil
}
