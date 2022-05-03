package domain

import "time"

type User struct {
	Id           string
	Firstname    string
	Lastname 	 string
	Email        string
	MobileNumber string
	Gender       string
	BirthDay     time.Time
	Username     string
	Biography    string
	Experience   string
	Education    string
	Skills       string
	Interests    string
	Password     string
	IsPrivate    bool
}