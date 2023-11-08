package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelx0liveira/message-oriented-middleware/cmd/api/config"
	"github.com/rafaelx0liveira/message-oriented-middleware/cmd/api/model"
	"github.com/rafaelx0liveira/message-oriented-middleware/cmd/api/util"
)

// Define a global variable to store the logger
var (
	logger *config.Logger
)

// Publish is used to publish a message
func PublishController (c *gin.Context) {
	request := model.Publish{}

	// Bind the request body to the request variable
	c.BindJSON(&request)

	// Initialize the logger
	logger = config.GetLogger("PublishController")

	// Validate the request
	if err := request.Validate(); err != nil {
		logger.Error(err.Error())
		util.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Printf("Content: %s\n", request)
}
