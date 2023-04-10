package hello_world_golang

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	// Create a new Gin router
	r := gin.Default()

	// Define a route handler for the "GET /" endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	return r
}
