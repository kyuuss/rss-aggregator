package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT must be set")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler())

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server started on port %s", port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
