package router

import (
	"consumer/api/controller"
	con "consumer/internal"

	"github.com/gin-gonic/gin"
)

// InitializeRoutes is used to initialize the routes
func InitializeRoutes(router *gin.Engine, consumer *con.Consumer) {
	v1 := router.Group("/api/v1")
	{
		// POST /api/v1/publish for publishing a message
		// The Publish function is from controllers package
		// Who inserts the Context in the Publish function is the Gin
		v1.POST("/publish", func(c *gin.Context) {
			c.Set("consumer", consumer)

			controller.PublishController(c)
		})
	}
}
