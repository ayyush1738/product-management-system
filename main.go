package main

import (
	"log"
	"net/http"

	"product-management-system/api"
	"product-management-system/db"
)

func main() {
	// Initialize the database connection
	db.InitDatabase()

	// Register API routes
	router := api.RegisterRoutes()

	// Start the server
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
