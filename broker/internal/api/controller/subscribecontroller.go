package controller

import (
	"github.com/gin-gonic/gin"
	"broker/internal/api/model"
	"broker/internal/api/config"
)

// Ticket is used to publish a webhook 
func SubscribeController (c *gin.Context) {

	// Define a variable to store the logger
	var logger *config.Logger

	// Initialize the message variable with the model.webhookdata struct
	message := model.WebhookData{}

	// Bind the request body to the message variable
	// This process (fill in a struct with the request body) is called "hydration"
	c.BindJSON(&message)

	// Initialize the logger
	logger = config.GetLogger("TicketController")

	// 
	
	// Log the successful message
	logger.Infof("Successfully published a webhook")
}