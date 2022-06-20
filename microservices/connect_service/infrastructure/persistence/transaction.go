package persistence

import (
	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (store *ConnectNeo4jDBStore) IsUserExistTx(tx neo4j.Transaction, userId primitive.ObjectID) (*bool, error) {
	query := "OPTIONAL MATCH (u:User{id:$userId}) RETURN u IS NOT NULL AS IsUserExist"
	params := map[string]interface{}{
		"userId": userId.Hex(),
	}
	result, err := tx.Run(query, params)
	if err != nil {
		return nil, err
	}
	record, err := result.Single()
	if err != nil {
		return nil, err
	}
	value, _ := record.Get("IsUserExist")
	exists := value.(bool)
	return &exists, err
}

func (store *ConnectNeo4jDBStore) IsConnectionExistTx(tx neo4j.Transaction, userId, cUserId primitive.ObjectID) (*bool, error) {
	query := "MATCH (u1:User{id:$userId}) MATCH (u2:User{id:$cUserId}) RETURN exists((u1)-[:Connection]-(u2)) AS Connection"
	params := map[string]interface{}{
		"userId":  userId.Hex(),
		"cUserId": cUserId.Hex(),
	}
	result, err := tx.Run(query, params)
	if err != nil {
		return nil, err
	}
	record, err := result.Single()
	if err != nil {
		return nil, err
	}
	value, _ := record.Get("Connection")
	exists := value.(bool)
	return &exists, err
}

func (store *ConnectNeo4jDBStore) IsReceivedInviteTx(tx neo4j.Transaction, userId, cUserId primitive.ObjectID) (*bool, error) {
	query := "MATCH (u1:User{id:$userId}) MATCH (u2:User{id:$cUserId}) RETURN exists((u1)<-[:Invite]-(u2)) AS IsReceived"
	params := map[string]interface{}{
		"userId":  userId.Hex(),
		"cUserId": cUserId.Hex(),
	}
	result, err := tx.Run(query, params)
	if err != nil {
		return nil, err
	}
	record, err := result.Single()
	if err != nil {
		return nil, err
	}
	value, _ := record.Get("IsReceived")
	exists := value.(bool)
	return &exists, err
}

func (store *ConnectNeo4jDBStore) IsInviteExistsTx(tx neo4j.Transaction, userId, cUserId primitive.ObjectID) (*bool, error) {
	query := "MATCH (u1:User{id:$userId}) MATCH (u2:User{id:$cUserId}) RETURN exists((u1)-[:Invite]->(u2)) AS IsInviteExist"
	params := map[string]interface{}{
		"userId":  userId.Hex(),
		"cUserId": cUserId.Hex(),
	}
	result, err := tx.Run(query, params)
	if err != nil {
		return nil, err
	}
	record, err := result.Single()
	if err != nil {
		return nil, err
	}
	value, _ := record.Get("IsInviteExist")
	exists := value.(bool)
	return &exists, err
}

func (store *ConnectNeo4jDBStore) PersistUserTx(tx neo4j.Transaction, user domain.Profile) (*domain.Profile, error) {
	query := "CREATE (u:User{id:$userId, name:$userId, private:$userPrivate})"
	params := map[string]interface{}{
		"userId":      user.Id.Hex(),
		"userPrivate": user.Private,
	}
	_, err := tx.Run(query, params)
	return nil, err
}

func (store *ConnectNeo4jDBStore) UpdateUserTx(tx neo4j.Transaction, user domain.Profile) (*domain.Profile, error) {
	query := "MATCH (u:User{id:$userId}) SET u.private = $userPrivate RETURN u"
	params := map[string]interface{}{
		"userId":      user.Id.Hex(),
		"userPrivate": user.Private,
	}
	_, err := tx.Run(query, params)
	return nil, err
}

func (store *ConnectNeo4jDBStore) IsUserPrivateTx(tx neo4j.Transaction, userId primitive.ObjectID) (*bool, error) {
	query := "MATCH (u:User{id:$userId}) RETURN u.private AS IsUserPrivate"
	params := map[string]interface{}{
		"userId": userId.Hex(),
	}
	result, err := tx.Run(query, params)
	if err != nil {
		return nil, err
	}
	record, err := result.Single()
	if err != nil {
		return nil, err
	}
	value, _ := record.Get("IsUserPrivate")
	private := value.(bool)
	return &private, err
}

func (store *ConnectNeo4jDBStore) CreateConnectionTx(tx neo4j.Transaction, userId, cUserId primitive.ObjectID) (*domain.Connection, error) {
	query := "MATCH (u1:User{id:$userId}) MATCH (u2:User{id:$cUserId}) MERGE (u1)-[c:Connection]-(u2)"
	params := map[string]interface{}{
		"userId":  userId.Hex(),
		"cUserId": cUserId.Hex(),
	}
	_, err := tx.Run(query, params)
	return nil, err
}

func (store *ConnectNeo4jDBStore) DeleteConnectionTx(tx neo4j.Transaction, userId, cUserId primitive.ObjectID) (*domain.Connection, error) {
	query := "MATCH (u1:User{id:$userId})-[c:Connection]-(u2:User{id:$cUserId}) DELETE c"
	params := map[string]interface{}{
		"userId":  userId.Hex(),
		"cUserId": cUserId.Hex(),
	}
	_, err := tx.Run(query, params)
	return nil, err
}

func (store *ConnectNeo4jDBStore) GetConnectionsTx(tx neo4j.Transaction, userId primitive.ObjectID) ([]*domain.Connection, error) {
	query := "MATCH (u:User{id:$userId}) MATCH (u)-[c:Connection]-(x) RETURN x.id AS UserId"
	params := map[string]interface{}{
		"userId": userId.Hex(),
	}
	result, err := tx.Run(query, params)
	if err != nil {
		return nil, err
	}
	var connections []*domain.Connection
	for result.Next() {
		id, _ := result.Record().Get("UserId")
		cUserId, err := primitive.ObjectIDFromHex(id.(string))
		if err != nil {
			return nil, err
		}
		connection := domain.Connection{
			User:  domain.Profile{Id: userId},
			CUser: domain.Profile{Id: cUserId},
		}
		connections = append(connections, &connection)
	}
	return connections, err
}

func (store *ConnectNeo4jDBStore) CreateInviteTx(tx neo4j.Transaction, userId, cUserId primitive.ObjectID) (*domain.Connection, error) {
	query := "MATCH (u1:User{id:$userId}) MATCH (u2:User{id:$cUserId}) MERGE (u1)-[i:Invite]->(u2)"
	params := map[string]interface{}{
		"userId":  userId.Hex(),
		"cUserId": cUserId.Hex(),
	}
	_, err := tx.Run(query, params)
	return nil, err
}

func (store *ConnectNeo4jDBStore) GetAllInvitationsTx(tx neo4j.Transaction, userId primitive.ObjectID) ([]*domain.Connection, error) {
	query := "MATCH (u:User{id:$userId}) MATCH (u)<-[i:Invite]-(x) RETURN x.id AS UserId"
	params := map[string]interface{}{
		"userId": userId.Hex(),
	}
	result, err := tx.Run(query, params)
	if err != nil {
		return nil, err
	}
	var invites []*domain.Connection
	for result.Next() {
		id, _ := result.Record().Get("UserId")
		cUserId, err := primitive.ObjectIDFromHex(id.(string))
		if err != nil {
			return nil, err
		}
		connection := domain.Connection{
			User:  domain.Profile{Id: userId},
			CUser: domain.Profile{Id: cUserId},
		}
		invites = append(invites, &connection)
	}
	return invites, err
}

func (store *ConnectNeo4jDBStore) GetAllSentInvitationsTx(tx neo4j.Transaction, userId primitive.ObjectID) ([]*domain.Connection, error) {
	query := "MATCH (u:User{id:$userId}) MATCH (u)-[i:Invite]->(x) RETURN x.id AS UserId"
	params := map[string]interface{}{
		"userId": userId.Hex(),
	}
	result, err := tx.Run(query, params)
	if err != nil {
		return nil, err
	}
	var invites []*domain.Connection
	for result.Next() {
		id, _ := result.Record().Get("UserId")
		cUserId, err := primitive.ObjectIDFromHex(id.(string))
		if err != nil {
			return nil, err
		}
		connection := domain.Connection{
			User:  domain.Profile{Id: userId},
			CUser: domain.Profile{Id: cUserId},
		}
		invites = append(invites, &connection)
	}
	return invites, err
}

func (store *ConnectNeo4jDBStore) DeleteInviteTx(tx neo4j.Transaction, userId, cUserId primitive.ObjectID) (*domain.Connection, error) {
	query := "MATCH (u1:User{id:$userId})-[i:Invite]->(u2:User{id:$cUserId}) DELETE i"
	params := map[string]interface{}{
		"userId":  userId.Hex(),
		"cUserId": cUserId.Hex(),
	}
	_, err := tx.Run(query, params)
	return nil, err
}

func (store *ConnectNeo4jDBStore) DeleteReceivedInviteTx(tx neo4j.Transaction, userId, cUserId primitive.ObjectID) (*domain.Connection, error) {
	query := "MATCH (u1:User{id:$userId})<-[i:Invite]-(u2:User{id:$cUserId}) DELETE i"
	params := map[string]interface{}{
		"userId":  userId.Hex(),
		"cUserId": cUserId.Hex(),
	}
	_, err := tx.Run(query, params)
	return nil, err
}
