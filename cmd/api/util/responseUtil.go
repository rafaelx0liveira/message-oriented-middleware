package util

import (
	"github.com/gin-gonic/gin"
)

func SendError(c *gin.Context, status int, message string) {
	c.Header("Content-Type", "application/json")
	c.JSON(status, gin.H{
		"error": message,
	})
}