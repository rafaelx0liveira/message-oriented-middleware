package subscriber

import (
	"broker/internal"
	"broker/internal/api/config"
	"broker/internal/api/model"
)

func Subscriber(webhookdata *model.WebhookData, logger *config.Logger, broker *internal.Broker) error {

	subscriber := internal.NewSubscriber(webhookdata.EventID, webhookdata.EventType, webhookdata.MessageData)

	err := broker.RegisterSubscriber(subscriber)

	if err != nil {
		logger.Errorf("Error while publishing message: %s", err.Error())
		return err
	}

	return nil
}
