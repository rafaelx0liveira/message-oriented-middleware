package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelx0liveira/message-oriented-middleware/internal/message_server/api/api/config"
	"github.com/rafaelx0liveira/message-oriented-middleware/internal/message_server/api/api/model"
	"github.com/rafaelx0liveira/message-oriented-middleware/internal/message_server/api/api/util"
	"github.com/rafaelx0liveira/message-oriented-middleware/internal/message_server/api/messagepublisher"
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
