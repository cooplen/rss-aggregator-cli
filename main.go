package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// creates a variable for the port defined in the .env
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("PORT is not found in the environment")
	}

	// creates routers
	rout := chi.NewRouter()

	// adds middleware to router to restrict usage; very wide permissions
	rout.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// new subrouter for /v1 namespace
	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)
	rout.Mount("/v1", v1Router)

	// creates server
	server := &http.Server{
		Handler: rout,
		Addr:    ":" + port,
	}

	log.Printf("Server starting on port %v", port)

	// starts server
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
