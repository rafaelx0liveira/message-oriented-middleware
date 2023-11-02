package consumer

type WebhookEvent struct {
	EventID   int
	EventType string
	Message   string
}
