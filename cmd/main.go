package main

import (
	"rodrigolluzdevr/api_rest/internal/db"
	"rodrigolluzdevr/api_rest/internal/routes"
)

func main() {
	// Connect mysql
	db.InitDB()

	// Routes API
	router := routes.SetupRoutes()

	router.Run(":8080")
}
