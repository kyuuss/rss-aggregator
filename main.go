package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/kyuuss/rss-aggregator/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT must be set")
	}

	db_url := os.Getenv("DB_URL")

	if db_url == "" {
		log.Fatal("DB_URL must be set")
	}

	connection, err := sql.Open("postgres", db_url)
	if err != nil {
		log.Fatal("Can not connect to database")
	}

	queries := database.New(connection)
	api_config := apiConfig{
		DB: queries,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1_router := chi.NewRouter()
	v1_router.Get("/health", handlerHealth)
	v1_router.Get("/error", handlerError)

	router.Mount("/v1", v1_router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server started on port %s", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
