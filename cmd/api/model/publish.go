package model

import (
	"github.com/rafaelx0liveira/message-oriented-middleware/cmd/api/util"
)

type Publish struct {
	Content string `json:"content"`
}

func (p *Publish) Validate () error {

	if p.Content == "" {
		return util.ErrParamIsRequired("Content", "string")
	}

	return nil
}