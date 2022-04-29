package persistence

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/domain"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

const (
	queryCreateConnection = "MERGE (user:User{id:$user, name:$user}) MERGE (u2:User{id:$userConnect, name:$userConnect}) MERGE (u1)-[c:connected]-(u2)"
)

type ConectionNeo4jDBStore struct {
	driver neo4j.Driver
}

func NewConectionNeo4jDBStore(driver neo4j.Driver) domain.ConnectionStore {
	return &ConectionNeo4jDBStore{driver: driver}
}

func (store *ConectionNeo4jDBStore) Connect(user, userConnect string) error {
	session, err := store.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	if err != nil {
		return err
	}
	defer session.Close()
	_, err = session.WriteTransaction(createConnectionTX(user, userConnect))
	if err != nil {
		return err
	}
	return nil
}

func createConnectionTX(user, userConnect string) neo4j.TransactionWork {
	return func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(queryCreateConnection, map[string]interface{}{"user": user, "userConnect": userConnect})
		if err != nil {
			return nil, err
		}
		return nil, result.Err()
	}
}

func (store *ConectionNeo4jDBStore) UnConnect() error {
	return nil
}
