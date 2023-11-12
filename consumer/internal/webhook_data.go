package consumer

type WebhookData struct {
	EventID     int    `json:"event_id"`
	EventType   string `json:"event_type"`
	MessageData string `json:"message_data"`
}
