package service

import (
	"net/http"

	"broker/internal/api/config"
	"broker/internal/api/publisher"
	"broker/internal/api/model"
	"broker/internal/api/util"
	"broker/internal"
	"github.com/gin-gonic/gin"
)

// Function to validate receiving a request
func ValidatePublishRequest(c *gin.Context, message *model.Message, logger *config.Logger, broker *internal.Broker) error{
	// Validate the request
	if err := message.Validate(); err != nil {
		logger.Error(err.Error())
		util.SendError(c, http.StatusBadRequest, err.Error())
		return err
	}

	err := publisher.PublishMessage(message, logger, broker)

	if err != nil {
		logger.Error(err.Error())
		util.SendError(c, http.StatusInternalServerError, err.Error())
		return err
	}

	return nil
}
