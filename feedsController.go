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

	// also create feed follow for the feed and user
	feedFollow, err := cfg.createFeedFollow(user.ID, feed.ID, r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respBody := struct {
		Feed database.Feed `json:"feed"`
		FeedFollow database.Feedfollow `json:"feed_follow"`
	}{
		Feed: feed,
		FeedFollow: feedFollow,
	}

	// Return created feed
	respondWithJSON(w, http.StatusOK, respBody)
}

func (cfg *apiConfig) getFeedsHandler(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, feeds)
}


func (cfg *apiConfig) postFeedFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	// decode the json request body
	type parameters struct {
		FeedId string `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		// handle decode parameters error 
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	feedId, err := uuid.Parse(params.FeedId)
	if err != nil {
		// handle decode parameters error 
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	feedFollow, err := cfg.createFeedFollow(user.ID, feedId, r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, feedFollow)
}


func(cfg *apiConfig) createFeedFollow(userID, feedID uuid.UUID, r *http.Request) (database.Feedfollow, error){
	newUUID := uuid.New()
	feedFollowParams := database.CreateFeedFollowParams{
		ID: newUUID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:userID,
		FeedID: feedID,
	}
	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), feedFollowParams)
	if err != nil {
		return database.Feedfollow{}, err
	}
	return feedFollow, nil
}