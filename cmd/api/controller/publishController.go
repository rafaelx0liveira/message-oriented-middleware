package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Publish is used to publish a message
func PublishController (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
