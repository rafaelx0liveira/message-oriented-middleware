package subscriber

import (
	"broker/internal"
	"broker/internal/api/config"
	"broker/internal/api/model"
)

func PublishTicket(message *model.WebhookData, logger *config.Logger) error{
	broker := internal.NewBroker()

	newMessage := internal.NewMessage(message.ID, message.Content)

	err := broker.SendMessage(newMessage)

	if err != nil {
		logger.Errorf("Error while publishing message: %s", err.Error())
		return err
	}

	return nil
}