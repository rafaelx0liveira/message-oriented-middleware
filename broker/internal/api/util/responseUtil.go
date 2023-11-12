package util

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendError(c *gin.Context, status int, message string) {
	c.Header("Content-Type", "application/json")
	c.JSON(status, gin.H{
		"error": message,
	})
}

func SendSuccess(c *gin.Context, operation string, data interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("operation %s occurred successfully", operation),
		"data":    data,
	})
}
