package handler

import (
	"encoding/json"

	"github.com/tockn/singo/manager"

	"github.com/gorilla/websocket"
)

type Handler struct {
	manager *manager.Manager
}

func NewHandler(man *manager.Manager) *Handler {
	return &Handler{manager: man}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type ReceiveMessageType string

var (
	ReceiveMessageTypeJoinRoom ReceiveMessageType = "join room"
	ReceiveMessageTypeOffer    ReceiveMessageType = "offer"
	ReceiveMessageTypeAnswer   ReceiveMessageType = "answer"
)

type ReceiveMessage struct {
	Type    ReceiveMessageType `json:"type"`
	Payload json.RawMessage    `json:"payload"`
}

type SendMessageType string

var (
	SendMessageTypeError  SendMessageType = "error"
	SendMessageTypeOffer  SendMessageType = "offer"
	SendMessageTypeAnswer SendMessageType = "answer"
)

type ErrorMessage string

func (e ErrorMessage) Error() string {
	return string(e)
}

var (
	ErrMsgInvalidType    ErrorMessage = "invalid type"
	ErrMsgInvalidPayload ErrorMessage = "invalid payload"
	ErrMsgInternalError  ErrorMessage = "internal error"
)

type SendMessage struct {
	Type    SendMessageType `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

func newErrorMessage(msg error) *SendMessage {
	return &SendMessage{
		Type:    SendMessageTypeError,
		Payload: []byte(msg.Error()),
	}
}
