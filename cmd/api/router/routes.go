package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitializeRoutes is used to initialize the routes
func InitializeRoutes(router *gin.Engine) {
	v1:=router.Group("/api/v1")
	{
		// POST /api/v1/publish for publishing a message
		v1.POST("/publish", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello World",
			})
		})

		// GET /api/v1/receive for receiving a message
		v1.GET("/receive", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello World",
			})
		})
	}

}

