package main

import (
	"github.com/gin-gonic/gin"
	"simple-app/cache"
	"simple-app/config"
	"simple-app/routes"
)

func main() {
	// Load configurations
	config.LoadConfig()

	// Connect to database
	config.ConnectDatabase()

	// Initialize Redis
	cache.InitRedis()

	// Set up router
	router := gin.Default()
	routes.RegisterRoutes(router)

	// Start the server
	router.Run(":8080")
}
