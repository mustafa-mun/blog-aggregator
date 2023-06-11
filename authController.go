package main

import (
	"errors"
	"net/http"
	"strings"

	"github.com/mustafa-mun/blog-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

// middleware for authentication of the user
func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user, err := cfg.AuthenticateUser(r)

		if err != nil {
			// Handle unauthorized access
			respondWithError(w, http.StatusUnauthorized, err.Error())
			
		} else {
		// Call the original handler function			
			handler(w, r, user)
		}
	}
}

func(cfg *apiConfig) AuthenticateUser (r *http.Request) (database.User, error){
	// get api key from auth header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
    // Handle the case when Authorization header is missing or empty
		return database.User{}, errors.New("missing api key")
	}
	apiKey := strings.Split(authHeader, " ")[1]

	user, err := cfg.DB.GetUser(r.Context(), apiKey)

	if err != nil {
		return database.User{}, errors.New("Unauthorized")
	}

	return user, nil
}
