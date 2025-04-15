package main

import (
	"github.com/gin-gonic/gin"
	"github.com/package/db"
	"github.com/package/routes"
)

func main() {
	// Setup MongoDB connection
	db.SetupDatabase()

	// Initialize Gin router
	r := gin.Default()

	r.Use(routes.Interceptor())

	// Setup routes
	routes.SetupRoutes(r)

	// Start the server
	r.Run(":8080")
}
