package service

import (
	"fmt"
	"net/http"

	"broker/internal"
	"broker/internal/api/config"
	"broker/internal/api/messagepublisher"
	"broker/internal/api/model"
	"broker/internal/api/util"

	"github.com/gin-gonic/gin"
)

// Function to validate receiving a request
func ValidateRequest(c *gin.Context, message *model.Message, logger *config.Logger, broker *internal.Broker) {
	// Validate the request
	if err := message.Validate(); err != nil {
		logger.Error(err.Error())
		util.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	messagepublisher.Init(broker)

	fmt.Printf("Request on SERVICE: %+v\n", *message)

}
