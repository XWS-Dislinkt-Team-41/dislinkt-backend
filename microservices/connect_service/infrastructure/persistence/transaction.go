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
			UserId:  userId,
			CUserId: cUserId,
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
			UserId:  userId,
			CUserId: cUserId,
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
			UserId:  userId,
			CUserId: cUserId,
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

func (store *ConnectNeo4jDBStore) DeleteAllInDBTx(tx neo4j.Transaction) error {
	query := "MATCH (n) DETACH DELETE n"
	_, err := tx.Run(query, map[string]interface{}{})
	return err
}

func (store *ConnectNeo4jDBStore) GetConnectionsOfUserConectionsTx(tx neo4j.Transaction, userId primitive.ObjectID) ([]*domain.Profile, error) {
	query := `MATCH (u:User{id:$userId}) MATCH (u)-[c:Connection]-(x)
		      WITH u AS self, collect(x) AS excluded
			  
		      MATCH (u1:User {id:$userId})-[:Connection*2..2]-(u2:User)
		      WITH self, excluded, collect(u2) AS suggestions, u2
		      WHERE NONE (u2 IN suggestions WHERE u2 IN excluded) AND u2 <> self
		      WITH DISTINCT u2
		      RETURN u2.id AS UserId, u2.private AS Private LIMIT 30`
	params := map[string]interface{}{
		"userId": userId.Hex(),
	}
	result, err := tx.Run(query, params)
	if err != nil {
		return nil, err
	}
	var users []*domain.Profile = []*domain.Profile{}
	for result.Next() {
		id, _ := result.Record().Get("UserId")
		p, _ := result.Record().Get("Private")
		userId, err := primitive.ObjectIDFromHex(id.(string))
		private := p.(bool)
		if err != nil {
			return nil, err
		}
		profile := domain.Profile{
			Id:      userId,
			Private: private,
		}
		users = append(users, &profile)
	}
	return users, err
}

func (store *ConnectNeo4jDBStore) LoadNodesFromCSVTx(tx neo4j.Transaction) error {
	query := `LOAD CSV WITH HEADERS FROM "file:///nodes.csv" AS row
			  WITH DISTINCT row 
			  MERGE (:User{id:row.id, name:row.name, private:(case row.private when 'true' then true else false end)})`
	_, err := tx.Run(query, map[string]interface{}{})
	return err
}

func (store *ConnectNeo4jDBStore) LoadRelationshipsFromCSVTx(tx neo4j.Transaction) error {
	query := `LOAD CSV WITH HEADERS FROM "file:///relationships.csv" AS row
			  WITH DISTINCT row 
			  MATCH (u1:User{id:row.userId}) MATCH (u2:User{id:row.cUserId}) MERGE (u1)-[c:Connection]-(u2)`
	_, err := tx.Run(query, map[string]interface{}{})
	return err
}
