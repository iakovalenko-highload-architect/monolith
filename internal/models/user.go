package models

import "time"

type User struct {
	ID         string
	Password   string
	FirstName  string
	SecondName string
	Birthday   time.Time
	City       string
	Biography  string
}

type Friendship struct {
	UserID   string
	FriendID string
}
