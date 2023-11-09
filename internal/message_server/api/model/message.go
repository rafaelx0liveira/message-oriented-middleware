package model

import (
	"github.com/google/uuid"
	"github.com/rafaelx0liveira/message-oriented-middleware/internal/message_server/api/api/util"
)

type Message struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func (m *Message) Validate () error {

	if m.Content == "" {
		return util.ErrParamIsRequired("Content", "string")
	}

	// generate a new ID for the message with uuid
	m.ID = uuid.New().String()

	return nil
}