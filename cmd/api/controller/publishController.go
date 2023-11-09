package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaelx0liveira/message-oriented-middleware/cmd/api/config"
	"github.com/rafaelx0liveira/message-oriented-middleware/cmd/api/model"
	"github.com/rafaelx0liveira/message-oriented-middleware/cmd/api/service"
)

// Define a global variable to store the logger
var (
	logger *config.Logger
)

// Publish is used to publish a message
func PublishController (c *gin.Context) {
	// Initialize the message variable with the model.Message struct
	message := model.Message{}

	// Bind the request body to the message variable
	// This process (fill in a struct with the request body) is called "hydration"
	c.BindJSON(&message)

	// Initialize the logger
	logger = config.GetLogger("PublishController")

	// Call the ValidateRequest function from service package
	service.ValidateRequest(c, &message, logger)

	// fmt.Printf("Request: %s\n", request)
}
