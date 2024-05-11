package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dotpep/rss-feed-aggr/internal/database"
	"github.com/google/uuid"
)

func (apiDBCfg *apiDBConfig) handlerCreateFeed(
	w http.ResponseWriter,
	r *http.Request,
	user database.User,
) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiDBCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}
