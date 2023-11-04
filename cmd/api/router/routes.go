package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaelx0liveira/message-oriented-middleware/cmd/api/controllers"
)

// InitializeRoutes is used to initialize the routes
func InitializeRoutes(router *gin.Engine) {
	v1:=router.Group("/api/v1")
	{
		// POST /api/v1/publish for publishing a message
		// The Publish function is from controllers package
		// Who inserts the Context in the Publish function is the Gin
		v1.POST("/publish", controllers.Publish)

		// GET /api/v1/receive for receiving a message
		// The Receive function is from controllers package
		// Who inserts the Context in the Receive function is the Gin
		v1.GET("/receive", controllers.Receive)
	}
}

