package main

import (
	"banktest/config"
	"banktest/routes"
)

func main() {
	// Initialize database connection
	config.InitDB()
	// defer config.DB.Close()

	// Initialize routes
	routes.Router()
}
