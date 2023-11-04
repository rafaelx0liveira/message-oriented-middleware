package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Publish is used to publish a message
func Publish (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

// Receive is used to receive a message
func Receive (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}