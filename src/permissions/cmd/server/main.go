package main

import (
	"fmt"
	"log"
	"net/http"

	"concierge/permissions/api"
	"concierge/permissions/internal/config"
)

func main() {
	fmt.Printf("Starting server")
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize router and default options
	// db, _ := database.NewMariaDB("root:default@tcp(127.0.0.1:3306)/permissions_db", true)
	// defer db.Close()
	// router := api.NewRouter(api.WithDB(db))

	// Start the server
	router := api.NewRouter()
	log.Printf("Starting server on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	//Local live
	// db, _ := database.NewMariaDB("root:default@tcp(127.0.0.1:3306)/permissions_db", true)
	// defer db.Close()
	// accCli := client.NewLiveAccountClient("http://localhost:8089")
	// router := api.NewRouter(api.WithDB(db), api.WithAccountClient(accCli))
	// if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
	// 	log.Fatalf("Error starting server: %v", err)
	// }
}
