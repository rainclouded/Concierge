package main

import (
	"log"
	"net/http"

	"concierge/permissions/api"
	"concierge/permissions/internal/config"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize router and default options
	router := api.NewRouter()

	// Start the server
	log.Printf("Starting server on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
