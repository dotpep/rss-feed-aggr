package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dotpep/rss-feed-aggr/internal/database"
	"github.com/google/uuid"
)

func (apiDBCfg *apiDBConfig) handlerCreateUser(resWriter http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Username string `json:"username"`
	}
	decoder := json.NewDecoder(req.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(resWriter, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiDBCfg.DB.CreateUser(req.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Username:  params.Username,
	})
	if err != nil {
		respondWithError(resWriter, 400, fmt.Sprintf(
			"Couldn't create user: %v, with username: %s", err, params.Username,
		))
	}

	respondWithJSON(resWriter, 200, databaseUserToUser(user))
}
