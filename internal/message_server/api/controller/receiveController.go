package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Receive is used to receive a message
func ReceiveController (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}