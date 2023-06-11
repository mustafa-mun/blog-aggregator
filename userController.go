package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mustafa-mun/blog-aggregator/internal/database"
)

func (cfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	// decode the json request body
	type parameters struct {
		Name string `json:"name"`
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
	userParams := database.CreateUserParams{
		ID: newUUID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: params.Name,
	}
	user, err := cfg.DB.CreateUser(r.Context(), userParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// Return created user 
	respondWithJSON(w, http.StatusOK, user)
}

func (cfg *apiConfig) getUserHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, user)
}