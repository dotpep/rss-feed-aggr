package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dotpep/rss-feed-aggr/internal/auth"
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
		return
	}

	respondWithJSON(resWriter, 201, databaseUserToUser(user))
}

func (apiDBCfg *apiDBConfig) handlerGetUserByAPIKey(resWriter http.ResponseWriter, req *http.Request) {
	apiKey, err := auth.GetAPIKey(req.Header)
	if err != nil {
		respondWithError(resWriter, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}

	user, err := apiDBCfg.DB.GetUserByAPIKey(req.Context(), apiKey)
	if err != nil {
		respondWithError(resWriter, 400, fmt.Sprintf("Couldn't get user: %v", err))
		return
	}

	respondWithJSON(resWriter, 200, databaseUserToUser(user))
}
