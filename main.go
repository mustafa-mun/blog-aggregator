package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	r := chi.NewRouter()
	subRouter := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{AllowedOrigins: []string{"*"}}))
	r.Mount("/v1", subRouter)
	
	subRouter.Get("/readiness", readinessHandler)
	subRouter.Get("/err", errHandler)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	server.ListenAndServe()
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	respBody := struct {
		Status string `json:"status"`
	} {
		Status: "ok",
	}
	respondWithJSON(w, http.StatusOK, respBody)	
}

func errHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}