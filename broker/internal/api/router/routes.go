package router

import (
	"broker/internal/api/controller"

	"github.com/gin-gonic/gin"
)

// InitializeRoutes is used to initialize the routes
func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		// POST /api/v1/publish for publishing a message
		// The Publish function is from controllers package
		// Who inserts the Context in the Publish function is the Gin
		v1.POST("/publish", controller.PublishController)

		// GET /api/v1/ticket for publishing a webhook
		// The Ticket function is from controllers package
		// Who inserts the Context in the Ticket function is the Gin
		v1.GET("/ticket", controller.TicketController)
	}
}
