package domain

import "time"

type Connection struct {
	Id          string    `bson:"id"`
	Timestamp   time.Time `bson:"timestamp"`
	User        Profile   `bson:"user"`
	UserConnect Profile   `bson:"userConnect"`
}

type Profile struct {
	Id string `bson:"id"`
}
