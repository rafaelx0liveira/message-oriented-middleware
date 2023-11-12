package model

import (
	"consumer/api/util"
)

type Message struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	Data string `json:"data"`
}

func (m *Message) Validate() error {

	if m.Type == "" {
		return util.ErrParamIsRequired("Type", "string")
	}
	if m.Data == "" {
		return util.ErrParamIsRequired("Data", "string")
	}
	return nil
}
