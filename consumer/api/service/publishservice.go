package service

import (
	"fmt"
	"net/http"

	"consumer/api/config"
	"consumer/api/model"
	"consumer/api/publisher"
	"consumer/api/util"
	con "consumer/internal"

	"github.com/gin-gonic/gin"
)

// Function to validate receiving a request
func ValidateRequest(c *gin.Context, message *model.Message, logger *config.Logger, consumer *con.Consumer) {
	// Validate the request
	if err := message.Validate(); err != nil {
		logger.Error(err.Error())
		util.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	publisher.PublishToConsumer(consumer, message)

	fmt.Printf("Request on SERVICE: %+v\n", *message)

}
