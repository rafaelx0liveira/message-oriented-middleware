package controller

import (
	"broker/internal"
	"broker/internal/api/config"
	"broker/internal/api/model"
	"broker/internal/api/service"
	"broker/internal/api/util"

	"github.com/gin-gonic/gin"
)

// Ticket is used to publish a webhook
func SubscribeController(c *gin.Context) {

	// Define a variable to store the logger
	var logger *config.Logger

	// Initialize the webhook variable with the model.webhookdata struct
	webhook := model.WebhookData{}

	// Bind the request body to the message variable
	// This process (fill in a struct with the request body) is called "hydration"
	c.BindJSON(&webhook)

	// Initialize the logger
	logger = config.GetLogger("SubscribeController")

	broker, exists := c.Get("broker")

	if exists {
		// Call the ValidateRequest function from service package
		err := service.ValidateSubscribeRequest(c, &webhook, logger, broker.(*internal.Broker))

		if err != nil {
			logger.Error(err.Error())
			return
		}
	} else {
		logger.Errorf("Error while publishing message: %s", "Broker not found")
		util.SendError(c, 500, "Broker not found")
		return
	}

	// Call the ValidateSubscribeRequest function from service package

	// Send a response to the client
	util.SendSuccess(c, "subscribe", webhook)

	// Log the successful message
	logger.Infof("Successfully published a webhook")
}
