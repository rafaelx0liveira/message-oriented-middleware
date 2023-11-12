package router

import (
	"broker/internal"
	"broker/internal/api/controller"

	"github.com/gin-gonic/gin"
)

// InitializeRoutes is used to initialize the routes
func InitializeRoutes(router *gin.Engine, broker *internal.Broker) {
	v1 := router.Group("/api/v1")
	{
		// POST /api/v1/publish for publishing a message
		// The Publish function is from controllers package
		// Who inserts the Context in the Publish function is the Gin
		v1.POST("/publish", func(c *gin.Context) {
			c.Set("broker", broker)

			controller.PublishController(c)
		})

		// GET /api/v1/subscribe for subscribing a webhook
		// The Subscribe function is from controllers package
		// Who inserts the Context in the Subscribe function is the Gin
		v1.POST("/subscribe", func(c *gin.Context) {
			c.Set("broker", broker)

			controller.SubscribeController(c)
		})
	}
}
