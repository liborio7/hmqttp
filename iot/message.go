package iot

import (
	"github.com/pkg/errors"
	"net/http"
)

var ErrEmptyTopic = errors.New("missing topic")
var ErrEmptyPayload = errors.New("missing payload")

type Message struct {
	Topic   string      `json:"topic"`
	Payload interface{} `json:"payload"`
}

func (m Message) Bind(req *http.Request) error {
	if m.Topic == "" {
		return ErrEmptyTopic
	}
	if m.Payload == nil {
		return ErrEmptyPayload
	}
	return nil
}
