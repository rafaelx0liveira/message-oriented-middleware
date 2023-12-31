package publisher

import (
	"broker/internal"
	"broker/internal/api/config"
	"broker/internal/api/model"
)

func PublishMessage(message *model.Message, logger *config.Logger, broker *internal.Broker) error{
	newMessage := internal.NewMessage(message.ID, message.Content)

	err := broker.SendMessage(newMessage)

	if err != nil {
		logger.Errorf("Error while publishing message: %s", err.Error())
		return err
	}

	return nil
}


