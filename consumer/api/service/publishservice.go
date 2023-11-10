package service

import (
	"fmt"
	"net/http"

	"consumer/api/config"
	"consumer/api/messagepublisher"
	"consumer/api/model"
	"consumer/api/util"

	"github.com/gin-gonic/gin"
)

// Function to validate receiving a request
func ValidateRequest(c *gin.Context, message *model.Message, logger *config.Logger) {
	// Validate the request
	if err := message.Validate(); err != nil {
		logger.Error(err.Error())
		util.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	messagepublisher.Init()

	fmt.Printf("Request on SERVICE: %+v\n", *message)

}
