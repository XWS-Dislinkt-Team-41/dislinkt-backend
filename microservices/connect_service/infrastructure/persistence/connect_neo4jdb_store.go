package persistence

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/domain"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	queryIsReceivedInvite      = "MATCH (u1:User{id:$userId}) MATCH (u2:User{id:$cUserId) RETURN exists((u1)<-[:Invite]-(u2)) AS IsReceived"
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
	session, err := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	if err != nil {
		return nil, err
	}
	defer session.Close()
	_, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(queryRegisterUser, map[string]interface{}{"userId": user.Id.Hex(), "userPrivate": user.Private})
		if err != nil {
			return nil, err
		}
		return result.Consume()
	})
	connection := domain.Profile{
		Id:      user.Id,
		Private: user.Private,
	}
	if err != nil {
		return nil, err
	}
	return &connection, nil
}

func (store *ConnectNeo4jDBStore) UpdateUser(user domain.Profile) (*domain.Profile, error) {
	session, err := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	if err != nil {
		return nil, err
	}
	defer session.Close()
	_, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(queryUpdateUser, map[string]interface{}{"userId": user.Id.Hex(), "userPrivate": user.Private})
		if err != nil {
			return nil, err
		}
		return result.Consume()
	})
	connection := domain.Profile{
		Id:      user.Id,
		Private: user.Private,
	}
	if err != nil {
		return nil, err
	}
	return &connection, nil
}

func (store *ConnectNeo4jDBStore) IsUserPrivate(userId primitive.ObjectID) (*bool, error) {
	session, err := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	if err != nil {
		return nil, err
	}
	defer session.Close()
	result, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(queryUpdateUser, map[string]interface{}{"userId": userId.Hex()})
		if err != nil {
			return nil, err
		}
		IsUserPrivate := true
		for result.Next() {
			if value, ok := result.Record().Get("IsUserPrivate"); ok {
				IsUserPrivate = value.(bool)
			} else {
				return nil, err
			}
		}
		return IsUserPrivate, result.Err()
	})
	IsUserPrivate := result.(bool)
	if err != nil {
		return nil, err
	}
	return &IsUserPrivate, nil
}

func (store *ConnectNeo4jDBStore) Connect(userId, cUserId primitive.ObjectID) (*domain.Connection, error) {
	session, err := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	if err != nil {
		return nil, err
	}
	defer session.Close()
	_, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(queryCreateConnection, map[string]interface{}{"userId": userId.Hex(), "cUserId": cUserId.Hex()})
		if err != nil {
			return nil, err
		}
		return result.Consume()
	})
	connection := domain.Connection{
		User:  domain.Profile{Id: userId},
		CUser: domain.Profile{Id: cUserId},
	}
	if err != nil {
		return nil, err
	}
	return &connection, nil
}

func (store *ConnectNeo4jDBStore) UnConnect(userId, cUserId primitive.ObjectID) error {
	session, err := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	if err != nil {
		return err
	}
	defer session.Close()
	_, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(queryDeleteConnection, map[string]interface{}{"userId": userId.Hex(), "cUserId": cUserId.Hex()})
		if err != nil {
			return nil, err
		}
		return result.Consume()
	})
	if err != nil {
		return err
	}
	return nil
}

func (store *ConnectNeo4jDBStore) GetUserConnections(userId primitive.ObjectID) ([]*domain.Connection, error) {
	session, err := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	if err != nil {
		return nil, err
	}
	defer session.Close()
	result, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(queryGetConnections, map[string]interface{}{"userId": userId.Hex()})
		if err != nil {
			return nil, err
		}
		var connections []*domain.Connection
		var id string
		for result.Next() {
			if value, ok := result.Record().GetByIndex(0).(string); ok {
				id = value
			} else {
				return nil, err
			}
			cUserId, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				return nil, err
			}
			connection := domain.Connection{
				User:  domain.Profile{Id: userId},
				CUser: domain.Profile{Id: cUserId},
			}
			connections = append(connections, &connection)
		}
		return connections, result.Err()
	})
	connections := result.([]*domain.Connection)
	if err != nil {
		return nil, err
	}
	return connections, nil
}

func (store *ConnectNeo4jDBStore) Invite(userId, cUserId primitive.ObjectID) (*domain.Connection, error) {
	session, err := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	if err != nil {
		return nil, err
	}
	defer session.Close()
	_, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(queryCreateInvite, map[string]interface{}{"userId": userId.Hex(), "cUserId": cUserId.Hex()})
		if err != nil {
			return nil, err
		}
		return nil, result.Err()
	})
	invite := domain.Connection{
		User:  domain.Profile{Id: userId},
		CUser: domain.Profile{Id: cUserId},
	}
	if err != nil {
		return nil, err
	}
	return &invite, nil
}

func (store *ConnectNeo4jDBStore) AcceptInvitation(userId primitive.ObjectID, cUserId primitive.ObjectID) (*domain.Connection, error) {
	session, err := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	if err != nil {
		return nil, err
	}
	defer session.Close()
	isCreatedConection := false
	_, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
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
	session, err := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	if err != nil {
		return err
	}
	defer session.Close()
	_, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
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
	session, err := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	if err != nil {
		return err
	}
	defer session.Close()
	_, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
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
	session, err := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	if err != nil {
		return nil, err
	}
	defer session.Close()
	result, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(queryGetAllInvitations, map[string]interface{}{"userId": userId.Hex()})
		if err != nil {
			return nil, err
		}

		var invites []*domain.Connection
		var id string
		for result.Next() {
			if value, ok := result.Record().GetByIndex(0).(string); ok {
				id = value
			} else {
				return nil, err
			}
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
	session, err := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	if err != nil {
		return nil, err
	}
	defer session.Close()
	result, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(queryGetAllSentInvitations, map[string]interface{}{"userId": userId.Hex()})
		if err != nil {
			return nil, err
		}

		var invites []*domain.Connection
		var id string
		for result.Next() {
			if value, ok := result.Record().GetByIndex(0).(string); ok {
				id = value
			} else {
				return nil, err
			}
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
