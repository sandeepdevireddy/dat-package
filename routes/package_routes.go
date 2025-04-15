package routes

import (
	"log"
	"time"

	"sandeepdevireddy/dat-package/controllers"

	"github.com/gin-gonic/gin"
)

func Interceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// before request
		c.Next()

		// after request
		latency := time.Since(t)
		log.Print("latency " + latency.String())

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

// SetupRoutes sets up all the routes for the company API ---
func SetupRoutes(r *gin.Engine) {
	r.POST("/packages", controllers.CreatePackages)
	r.POST("/package/components", controllers.CreatePackageComponents)

}
