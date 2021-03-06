package persistence

import (
	"fmt"

	"github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/connect_service/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ConnectNeo4jDBStore struct {
	driver *neo4j.Driver
}

func NewConnectNeo4jDBStore(driver *neo4j.Driver) domain.ConnectStore {
	return &ConnectNeo4jDBStore{driver: driver}
}

func (store *ConnectNeo4jDBStore) CheckIfUserDoesNotExist(tx neo4j.Transaction, userId primitive.ObjectID) (*bool, error) {
	userExists, err := store.IsUserExistTx(tx, userId)
	if err != nil {
		return nil, err
	}
	if *userExists {
		err = status.Error(codes.AlreadyExists, fmt.Sprintf("user: %s already exists", userId.Hex()))
		return nil, err
	}
	return userExists, nil
}

func (store *ConnectNeo4jDBStore) CheckIfUserExists(tx neo4j.Transaction, userId primitive.ObjectID) (*bool, error) {
	userExists, err := store.IsUserExistTx(tx, userId)
	if err != nil {
		return nil, err
	}
	if !*userExists {
		err = status.Error(codes.InvalidArgument, fmt.Sprintf("user: %s does not exist", userId.Hex()))
		return nil, err
	}
	return userExists, nil
}

func (store *ConnectNeo4jDBStore) CheckIfConnectionDoesNotExist(tx neo4j.Transaction, userId, cUserId primitive.ObjectID) (*bool, error) {
	connectionExists, err := store.IsConnectionExistTx(tx, userId, cUserId)
	if err != nil {
		return nil, err
	}
	if *connectionExists {
		err = status.Error(codes.InvalidArgument, "connection already exists")
		return nil, err
	}
	return connectionExists, nil
}

func (store *ConnectNeo4jDBStore) CheckIfReceivedInviteExist(tx neo4j.Transaction, userId, cUserId primitive.ObjectID) (*bool, error) {
	inviteExists, err := store.IsReceivedInviteTx(tx, userId, cUserId)
	if err != nil {
		return nil, err
	}
	if !*inviteExists {
		err = status.Error(codes.InvalidArgument, "received invite does not exists")
		return nil, err
	}
	return inviteExists, nil
}

func (store *ConnectNeo4jDBStore) CheckIfInviteExist(tx neo4j.Transaction, userId, cUserId primitive.ObjectID) (*bool, error) {
	inviteExists, err := store.IsInviteExistsTx(tx, userId, cUserId)
	if err != nil {
		return nil, err
	}
	if !*inviteExists {
		err = status.Error(codes.InvalidArgument, "invite does not exists")
		return nil, err
	}
	return inviteExists, nil
}

func (store *ConnectNeo4jDBStore) CheckIfNotBlocked(tx neo4j.Transaction, userId, bUserId primitive.ObjectID) (*bool, error) {
	blockExists, err := store.IsBlockExistsTx(tx, userId, bUserId)
	if err != nil {
		return nil, err
	}
	if *blockExists {
		err = status.Error(codes.InvalidArgument, "user is blocked or you are blocked by user")
		return nil, err
	}
	return blockExists, nil
}

func (store *ConnectNeo4jDBStore) Register(user domain.Profile) (*domain.Profile, error) {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := store.CheckIfUserDoesNotExist(tx, user.Id)
		if err != nil {
			return nil, err
		}
		_, err = store.PersistUserTx(tx, user)
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
		_, err := store.CheckIfUserExists(tx, user.Id)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		_, err = store.UpdateUserTx(tx, user)
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
		_, err := store.CheckIfUserExists(tx, userId)
		if err != nil {
			return nil, err
		}
		IsUserPrivate, err = store.IsUserPrivateTx(tx, userId)
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
		_, err := store.CheckIfUserExists(tx, userId)
		if err != nil {
			return nil, err
		}
		_, err = store.CheckIfUserExists(tx, cUserId)
		if err != nil {
			return nil, err
		}
		_, err = store.CheckIfNotBlocked(tx, userId, cUserId)
		if err != nil {
			return nil, err
		}
		_, err = store.CreateConnectionTx(tx, userId, cUserId)
		return nil, err
	})
	if err != nil {
		return nil, err
	}
	connection := domain.Connection{
		UserId:  userId,
		CUserId: cUserId,
	}
	return &connection, nil
}

func (store *ConnectNeo4jDBStore) UnConnect(userId, cUserId primitive.ObjectID) error {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := store.CheckIfUserExists(tx, userId)
		if err != nil {
			return nil, err
		}
		_, err = store.CheckIfUserExists(tx, cUserId)
		if err != nil {
			return nil, err
		}
		_, err = store.DeleteConnectionTx(tx, userId, cUserId)
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
		_, err := store.CheckIfUserExists(tx, userId)
		if err != nil {
			return nil, err
		}
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
		_, err := store.CheckIfUserExists(tx, userId)
		if err != nil {
			return nil, err
		}
		_, err = store.CheckIfUserExists(tx, cUserId)
		if err != nil {
			return nil, err
		}
		_, err = store.CheckIfNotBlocked(tx, userId, cUserId)
		if err != nil {
			return nil, err
		}
		_, err = store.CheckIfConnectionDoesNotExist(tx, userId, cUserId)
		if err != nil {
			return nil, err
		}
		_, err = store.CreateInviteTx(tx, userId, cUserId)
		return nil, err
	})
	if err != nil {
		return nil, err
	}
	invite := domain.Connection{
		UserId:  userId,
		CUserId: cUserId,
	}
	return &invite, nil
}

func (store *ConnectNeo4jDBStore) AcceptInvitation(userId primitive.ObjectID, cUserId primitive.ObjectID) (*domain.Connection, error) {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	isCreatedConection := false
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := store.CheckIfReceivedInviteExist(tx, userId, cUserId)
		if err != nil {
			return nil, err
		}
		_, err = store.DeleteReceivedInviteTx(tx, userId, cUserId)
		if err != nil {
			return nil, err
		}
		_, err = store.CreateConnectionTx(tx, userId, cUserId)
		if err != nil {
			return nil, err
		}
		isCreatedConection = true
		return nil, err
	})
	if err != nil {
		return nil, err
	}
	var connection *domain.Connection
	if isCreatedConection {
		connection = &domain.Connection{
			UserId:  userId,
			CUserId: cUserId,
		}
	}
	return connection, nil
}

func (store *ConnectNeo4jDBStore) DeclineInvitation(userId primitive.ObjectID, cUserId primitive.ObjectID) error {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := store.CheckIfReceivedInviteExist(tx, userId, cUserId)
		if err != nil {
			return nil, err
		}
		_, err = store.DeleteReceivedInviteTx(tx, userId, cUserId)
		if err != nil {
			return nil, err
		}
		return nil, err
	})
	if err != nil {
		return err
	}
	return nil
}

func (store *ConnectNeo4jDBStore) CancelInvitation(userId primitive.ObjectID, cUserId primitive.ObjectID) error {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := store.CheckIfInviteExist(tx, userId, cUserId)
		if err != nil {
			return nil, err
		}
		_, err = store.DeleteInviteTx(tx, userId, cUserId)
		if err != nil {
			return nil, err
		}
		return nil, err
	})
	if err != nil {
		return err
	}
	return nil
}

func (store *ConnectNeo4jDBStore) GetAllInvitations(userId primitive.ObjectID) ([]*domain.Connection, error) {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := store.CheckIfUserExists(tx, userId)
		if err != nil {
			return nil, err
		}
		result, err := store.GetAllInvitationsTx(tx, userId)
		return result, err
	})
	if err != nil {
		return nil, err
	}
	invites := result.([]*domain.Connection)
	return invites, nil
}

func (store *ConnectNeo4jDBStore) GetAllSentInvitations(userId primitive.ObjectID) ([]*domain.Connection, error) {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := store.CheckIfUserExists(tx, userId)
		if err != nil {
			return nil, err
		}
		result, err := store.GetAllSentInvitationsTx(tx, userId)
		return result, err
	})
	if err != nil {
		return nil, err
	}
	invites := result.([]*domain.Connection)
	return invites, nil
}

func (store *ConnectNeo4jDBStore) InitNeo4jDB() error {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		err := store.DeleteAllInDBTx(tx)
		if err != nil {
			return nil, err
		}
		err = store.LoadNodesFromCSVTx(tx)
		if err != nil {
			return nil, err
		}
		err = store.LoadRelationshipsFromCSVTx(tx)
		return nil, err
	})
	if err != nil {
		return err
	}
	return nil
}

func (store *ConnectNeo4jDBStore) GetUserSuggestions(userId primitive.ObjectID) ([]*domain.Profile, error) {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := store.CheckIfUserExists(tx, userId)
		if err != nil {
			return nil, err
		}
		result, err := store.GetConnectionsOfUserConectionsTx(tx, userId)
		return result, err
	})
	if err != nil {
		return nil, err
	}
	users := result.([]*domain.Profile)
	return users, nil
}

func (store *ConnectNeo4jDBStore) GetRandomUsersWithoutConections(userId primitive.ObjectID) ([]*domain.Profile, error) {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := store.CheckIfUserExists(tx, userId)
		if err != nil {
			return nil, err
		}
		result, err := store.GetRandomUsersWithoutConectionsTx(tx, userId)
		return result, err
	})
	if err != nil {
		return nil, err
	}
	users := result.([]*domain.Profile)
	return users, nil
}

func (store *ConnectNeo4jDBStore) GetRandomUsers(userId primitive.ObjectID) ([]*domain.Profile, error) {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := store.CheckIfUserExists(tx, userId)
		if err != nil {
			return nil, err
		}
		result, err := store.GetRandomUsersTx(tx, userId)
		return result, err
	})
	if err != nil {
		return nil, err
	}
	users := result.([]*domain.Profile)
	return users, nil
}

func (store *ConnectNeo4jDBStore) Block(userId primitive.ObjectID, bUserId primitive.ObjectID) (*domain.Block, error) {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := store.CheckIfUserExists(tx, userId)
		if err != nil {
			return nil, err
		}
		_, err = store.CheckIfUserExists(tx, bUserId)
		if err != nil {
			return nil, err
		}
		_, err = store.CreateBlockTx(tx, userId, bUserId)
		if err != nil {
			return nil, err
		}
		_, err = store.DeleteConnectionTx(tx, userId, bUserId)
		if err != nil {
			return nil, err
		}
		_, err = store.DeleteReceivedInviteTx(tx, userId, bUserId)
		if err != nil {
			return nil, err
		}
		_, err = store.DeleteInviteTx(tx, userId, bUserId)
		return nil, err
	})
	if err != nil {
		return nil, err
	}
	block := domain.Block{
		UserId:  userId,
		BUserId: bUserId,
	}
	return &block, nil
}

func (store *ConnectNeo4jDBStore) GetBlockedUsers(userId primitive.ObjectID) ([]*domain.Block, error) {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := store.CheckIfUserExists(tx, userId)
		if err != nil {
			return nil, err
		}
		result, err := store.GetBlockedUsersTx(tx, userId)
		return result, err
	})
	if err != nil {
		return nil, err
	}
	blocks := result.([]*domain.Block)
	return blocks, nil
}

func (store *ConnectNeo4jDBStore) UnBlock(userId primitive.ObjectID, bUserId primitive.ObjectID) error {
	session := (*store.driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := store.CheckIfUserExists(tx, userId)
		if err != nil {
			return nil, err
		}
		_, err = store.CheckIfUserExists(tx, bUserId)
		if err != nil {
			return nil, err
		}
		_, err = store.DeleteBlockTx(tx, userId, bUserId)
		return nil, err
	})
	if err != nil {
		return err
	}
	return nil
}
