package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ysdinesh31/temperature-error-logs/internal/db"
	"github.com/ysdinesh31/temperature-error-logs/internal/routes"
)

func main() {
	// Initialize MongoDB
	if err := db.ConnectMongoDB(); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	r := mux.NewRouter()

	routes.SetupRoutes(r)

	// Start server
	log.Print("Starting server on :8080")
	port := os.Getenv("PORT")
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
