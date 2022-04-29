package persistence

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func GetClient(host, username, password, port string) (neo4j.Driver, error) {
	uri := fmt.Sprintf("neo4j://%s:%s/", host, port)
	return neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
}
