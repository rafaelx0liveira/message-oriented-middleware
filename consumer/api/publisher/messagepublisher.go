package publisher

import (
	"consumer/api/model"
	con "consumer/internal"
)

func PublishToConsumer(consumer *con.Consumer, message *model.Message) {
	consumer.HookCB(&con.WebhookData{
		EventID:     message.ID,
		EventType:   message.Type,
		MessageData: message.Data,
	})
}
