package main

import (
	"encoding/json"
	"net/http"
	"strings"
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

func (cfg *apiConfig) getUserHandler(w http.ResponseWriter, r *http.Request) {
	// get api key from auth header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
    // Handle the case when Authorization header is missing or empty
		respondWithError(w, http.StatusUnauthorized, "missing api key")
		return
	}
	apiKey := strings.Split(authHeader, " ")[1]

	user, err := cfg.DB.GetUser(r.Context(), apiKey)

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	// user is found
	respondWithJSON(w, http.StatusOK, user)
}