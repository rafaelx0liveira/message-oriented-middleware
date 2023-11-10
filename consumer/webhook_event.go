package consumer

import "time"

type WebhookEvent struct {
	PostData  *WebhookData
	Timeout   int
	StartTime time.Time
}
