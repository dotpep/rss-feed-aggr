package main

import (
	"time"

	"github.com/dotpep/rss-feed-aggr/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	APIKey    string    `json:"api_key"`
}

// TODO: Make Out and In structs and responses with dbEntityToEntity
type UserOut struct {
	Username string `json:"username"`
}

func databaseUserToUser(dbUser database.User) User {
	//return UserOut{
	//	Username: dbUser.Username,
	//}
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Username:  dbUser.Username,
		APIKey:    dbUser.ApiKey,
	}
}
