package client

import (
	"fmt"
	"net/http"
)

type Message struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Data string `json:"data"`
}

func (m *Message) Validate() error {
	if m.ID == "" {
		return fmt.Errorf("id is required")
	}
	if m.Type == "" {
		return fmt.Errorf("type is required")
	}
	if m.Data == "" {
		return fmt.Errorf("data is required")
	}
	return nil
}

func PostMessage(url string, message *Message) error {
	// Validate the request
	if err := message.Validate(); err != nil {
		return err
	}

	// Send the request
	_, err := http.Post(url, "application/json", nil)

	if err != nil {
		return err
	}

	return nil
}