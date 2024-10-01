package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/rahiyansafz/go-mongo-todos/db"
	"github.com/rahiyansafz/go-mongo-todos/handlers"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	if err := db.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	http.HandleFunc("/todos", handlers.TodoHandler)
	http.HandleFunc("/todos/", handlers.TodoHandler)
	http.HandleFunc("/search", handlers.SearchHandler)

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
