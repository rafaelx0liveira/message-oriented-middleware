package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelx0liveira/message-oriented-middleware/cmd/api/config"
	"github.com/rafaelx0liveira/message-oriented-middleware/cmd/api/model"
	"github.com/rafaelx0liveira/message-oriented-middleware/cmd/api/util"
)

// Function to validate receiving a request
func ValidateRequest(c *gin.Context, request model.Publish, logger *config.Logger) {
	// Validate the request
	if err := request.Validate(); err != nil {
		logger.Error(err.Error())
		util.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Printf("Request on SERVICE: %s\n", request)
}
