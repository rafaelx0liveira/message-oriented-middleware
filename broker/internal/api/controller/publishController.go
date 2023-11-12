package controller

import (
	"broker/internal"
	"broker/internal/api/config"
	"broker/internal/api/model"
	"broker/internal/api/service"
	"broker/internal/api/util"

	"github.com/gin-gonic/gin"
)

// Publish is used to publish a message
func PublishController(c *gin.Context) {

	// Define a variable to store the logger
	var logger *config.Logger

	// Initialize the message variable with the model.Message struct
	message := model.Message{}

	// Bind the request body to the message variable
	// This process (fill in a struct with the request body) is called "hydration"
	c.BindJSON(&message)

	// Initialize the logger
	logger = config.GetLogger("PublishController")

	broker, exists := c.Get("broker")

	if exists {
		// Call the ValidateRequest function from service package
		service.ValidatePublishRequest(c, &message, logger, broker.(*internal.Broker))
	} else {
		logger.Errorf("Error while publishing message: %s", "Broker not found")
		util.SendError(c, 500, "Broker not found")
		return
	}

	// Send a response to the client
	util.SendSuccess(c, "publish message", message)

	// Log the successful message
	logger.Infof("Successfully published message: {ID: %v, Content: %v}", message.ID, message.Content)
}
