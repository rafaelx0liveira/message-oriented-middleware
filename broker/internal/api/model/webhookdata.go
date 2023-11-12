package model

import (
	"broker/internal/api/util"
	"math/rand"
	"time"
)

type WebhookData struct {
	EventID     int 	 `json:"event_id"`
	EventType   string `json:"event_type"`
	MessageData string `json:"message_data"`
}

func (w *WebhookData) Validate() error {
	if w.MessageData == "" {
		return util.ErrParamIsRequired("MessageData", "string")
	}

	if w.EventType == "" {
		return util.ErrParamIsRequired("EventType", "string")
	}

	// seed to get different result every time
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// generate a new random ID for the message
	w.EventID = rand.Intn(100000)
	
	return nil
}