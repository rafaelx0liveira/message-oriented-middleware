package model

import (
	"broker/internal/api/util"
	"github.com/google/uuid"
)

type Message struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func (m *Message) Validate() error {

	if m.Content == "" {
		return util.ErrParamIsRequired("Content", "string")
	}

	// generate a new ID for the message with uuid
	m.ID = uuid.New().String()

	return nil
}
