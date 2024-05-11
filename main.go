package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/dotpep/rss-feed-aggr/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiDBConfig struct {
	DB *database.Queries
}

func main() {
	// Environment Variable
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	// Database
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	apiDBCfg := apiDBConfig{
		DB: database.New(conn),
	}

	// Routers (Endpoints)
	router := chi.NewRouter()

	// cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// v1 router endpoints
	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)

	v1Router.Post("/users", apiDBCfg.handlerCreateUser)
	v1Router.Get("/users", apiDBCfg.middlewareAuth(apiDBCfg.handlerGetUserByAPIKey))

	v1Router.Post("/feeds", apiDBCfg.middlewareAuth(apiDBCfg.handlerCreateFeed))

	router.Mount("/v1", v1Router)

	// Http Server (JSON REST API)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
