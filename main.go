package main

import (
	"log"

	"sharingvision-backend/config"
	"sharingvision-backend/routes"
)

func main() {
	// Connect to database & AutoMigrate
	config.ConnectDatabase()

	// Setup router
	r := routes.SetupRouter()

	// Start server
	log.Println("Server running at http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
