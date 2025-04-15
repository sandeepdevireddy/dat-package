package main

import (
	"sandeepdevireddy/dat-package/db"

	"sandeepdevireddy/dat-package/routes"

	"github.com/gin-gonic/gin"
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
