package persistence

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/domain"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

const (
	queryCreateConnection = "MERGE (u1:User{id:$user, name:$user}) MERGE (u2:User{id:$userConnect, name:$userConnect}) MERGE (u1)-[c:connection]-(u2)"
	queryDeleteConnection = "MATCH (u1:User{id:$user, name:$user})-[c:connection]-(u2:User{id:$userConnect, name:$userConnect}) DELETE c"
	queryGetConnections   = "MATCH (u:User{id:$user}) MATCH (u)-[c:connection]->(x) RETURN x.id"
)

type ConnectNeo4jDBStore struct {
	driver *neo4j.Driver
}

func NewConnectNeo4jDBStore(driver *neo4j.Driver) domain.ConnectStore {
	return &ConnectNeo4jDBStore{driver: driver}
}

func (store *ConnectNeo4jDBStore) Connect(user, userConnect string) error {
	session, err := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	if err != nil {
		return err
	}
	defer session.Close()
	_, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(queryCreateConnection, map[string]interface{}{"user": user, "userConnect": userConnect})
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

func (store *ConnectNeo4jDBStore) UnConnect(user, userConnect string) error {
	session, err := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	if err != nil {
		return err
	}
	defer session.Close()
	_, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(queryDeleteConnection, map[string]interface{}{"user": user, "userConnect": userConnect})
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

func (store *ConnectNeo4jDBStore) GetUserConnections(user string) ([]string, error) {
	session, err := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	if err != nil {
		return nil, err
	}
	defer session.Close()
	connections, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(queryGetConnections, map[string]interface{}{"user": user})
		if err != nil {
			return nil, err
		}

		var usersIds []string
		for result.Next() {
			if id, ok := result.Record().GetByIndex(0).(string); ok {
				usersIds = append(usersIds, id)
			}
		}

		return usersIds, result.Err()
	})

	if err != nil {
		return nil, err
	}
	return connections.([]string), nil
}
