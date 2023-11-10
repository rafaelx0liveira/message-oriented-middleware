package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"message_server/api/config"
	"message_server/api/model"
	"message_server/api/util"
	"message_server/api/messagepublisher"
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
