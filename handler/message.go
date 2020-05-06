package handler

import (
	"encoding/json"
	"log"

	"github.com/tockn/singo/model"
)

type ReceiveMessageType string

var (
	ReceiveMessageTypeJoinRoom ReceiveMessageType = "join"
	ReceiveMessageTypeOffer    ReceiveMessageType = "offer"
	ReceiveMessageTypeAnswer   ReceiveMessageType = "answer"
)

type ReceiveMessage struct {
	Type    ReceiveMessageType `json:"type"`
	Payload json.RawMessage    `json:"payload"`
}

type ErrorMessage string

func (e ErrorMessage) Error() string {
	return string(e)
}

var (
	ErrMsgInvalidType    ErrorMessage = "invalid type"
	ErrMsgInvalidPayload ErrorMessage = "invalid payload"
	ErrMsgInternalError  ErrorMessage = "internal error"
)

func newErrorMessage(msg error) *model.Message {
	log.Println("error:", msg)
	return &model.Message{
		Type:    model.MessageTypeError,
		Payload: []byte(msg.Error()),
	}
}
