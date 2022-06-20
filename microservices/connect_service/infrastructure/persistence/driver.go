package persistence

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetDriver(host, username, password, port string) (*neo4j.Driver, error) {
	uri := fmt.Sprintf("bolt://%s:%s/", host, port)
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, err
	}
	return &driver, nil
}
