package controller

import (
	"broker/internal"
	"broker/internal/api/config"
	"broker/internal/api/model"
	"broker/internal/api/service"
	"github.com/gin-gonic/gin"
)

// Define a global variable to store the logger
var (
	logger *config.Logger
)

// Publish is used to publish a message
func PublishController(c *gin.Context) {
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
		service.ValidateRequest(c, &message, logger, broker.(*internal.Broker))

		// fmt.Printf("Request: %s\n", request)
	}
}
