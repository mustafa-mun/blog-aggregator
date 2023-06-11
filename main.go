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

type AuthHandler interface {
	func(http.ResponseWriter, *http.Request, database.User)
}

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
	apiRouter := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{AllowedOrigins: []string{"*"}}))
	r.Mount("/v1", apiRouter)
	
	apiRouter.Get("/readiness", readinessHandler)
	apiRouter.Get("/err", errHandler)
	apiRouter.Get("/users", apiCfg.middlewareAuth(apiCfg.getUserHandler)) // get current user (Auth ApiKey Route)
	apiRouter.Post("/users", apiCfg.createUserHandler)
	apiRouter.Post("/feeds", apiCfg.middlewareAuth(apiCfg.postFeedHandler))

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






