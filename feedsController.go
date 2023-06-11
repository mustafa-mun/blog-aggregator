package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mustafa-mun/blog-aggregator/internal/database"
)

func (cfg *apiConfig) postFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	// decode the json request body
	type parameters struct {
		Name string `json:"name"`
		Url string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		// handle decode parameters error 
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	newUUID := uuid.New()
	feedParams := database.CreateFeedParams{
		ID: newUUID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: params.Name,
		Url: params.Url,
		UserID: user.ID,
	}
	
	feed, err := cfg.DB.CreateFeed(r.Context(), feedParams)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// Return created feed
	respondWithJSON(w, http.StatusOK, feed)
}

func (cfg *apiConfig) getFeedsHandler(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, feeds)
}