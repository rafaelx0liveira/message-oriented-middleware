package service

import (
	"net/http"

	"broker/internal"
	"broker/internal/api/config"
	"broker/internal/api/model"
	"broker/internal/api/subscriber"
	"broker/internal/api/util"

	"github.com/gin-gonic/gin"
)

// Function to validate receiving a request
func ValidateSubscribeRequest(c *gin.Context, webhookdata *model.WebhookData, logger *config.Logger, broker *internal.Broker) error {
	// Validate the request
	if err := webhookdata.Validate(); err != nil {
		logger.Error(err.Error())
		util.SendError(c, http.StatusBadRequest, err.Error())
		return err
	}

	err := subscriber.Subscriber(webhookdata, logger, broker)

	if err != nil {
		logger.Error(err.Error())
		util.SendError(c, http.StatusInternalServerError, err.Error())
		return err
	}

	return nil
}
