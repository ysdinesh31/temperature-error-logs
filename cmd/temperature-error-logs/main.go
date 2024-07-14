package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/ysdinesh31/temperature-error-logs/internal/db"
	"github.com/ysdinesh31/temperature-error-logs/internal/routes"
)

func main() {
	// Initialize MongoDB
	if err := db.ConnectMongoDB(); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	r := mux.NewRouter()

	// Setup routes
	routes.SetupRoutes(r)

	// Start server
	logrus.Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logrus.Fatalf("Failed to start server: %v", err)
	}
}
