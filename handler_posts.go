package main

import (
	"fmt"
	"net/http"

	"github.com/dotpep/rss-feed-aggr/internal/database"
)

func (apiDBCfg *apiDBConfig) handlerGetPostListForUser(
	w http.ResponseWriter,
	r *http.Request,
	user database.User,
) {
	postList, err := apiDBCfg.DB.GetPostListForUser(r.Context(), database.GetPostListForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get post list: %v", err))
		return
	}

	respondWithJSON(w, 200, databasePostListToPostList(postList))
}
