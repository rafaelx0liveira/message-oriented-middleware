package internal

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type WebhookData struct {
	EventID     int    `json:"id"`
	EventType   string `json:"type"`
	MessageData string `json:"data"`
}

func PostMessage(subs *Subscriber, message *Message) error {

	payload := WebhookData{
		EventID:     subs.ID,
		EventType:   subs.EventType,
		MessageData: message.Content,
	}

	jsonData, errJsn := json.Marshal(payload)
	if errJsn != nil {
		panic(errJsn)
	}

	// Send the request
	_, err := http.Post(subs.ConsumerUrl, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		return err
	}

	return nil
}
