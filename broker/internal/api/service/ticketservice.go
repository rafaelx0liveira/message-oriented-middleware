package service

import (
	"net/http"

	"broker/internal/api/config"
	"broker/internal/api/messagepublisher"
	"broker/internal/api/model"
	"broker/internal/api/util"

	"github.com/gin-gonic/gin"
)

// Function to validate receiving a request
func ValidateTicketRequest(c *gin.Context, message *model.WebhookData, logger *config.Logger) error{
	// Validate the request
	if err := message.Validate(); err != nil {
		logger.Error(err.Error())
		util.SendError(c, http.StatusBadRequest, err.Error())
		return err
	}

	err := messagepublisher.PublishTicket(message, logger)

	if err != nil {
		logger.Error(err.Error())
		util.SendError(c, http.StatusInternalServerError, err.Error())
		return err
	}

	return nil
}