package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mustafa-mun/blog-aggregator/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("CONN")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	apiCfg := apiConfig{DB: dbQueries}

	port := os.Getenv("PORT")
	r := chi.NewRouter()
	subRouter := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{AllowedOrigins: []string{"*"}}))
	r.Mount("/v1", subRouter)
	
	subRouter.Get("/readiness", readinessHandler)
	subRouter.Get("/err", errHandler)
	subRouter.Post("/users", apiCfg.createUserHandler)

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



