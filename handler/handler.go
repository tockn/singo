package handler

import (
	"encoding/json"
	"log"
	"net/http"

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
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

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

type SendMessageType string

var (
	SendMessageTypeNotifyClientID SendMessageType = "notify-client-id"
	SendMessageTypeError          SendMessageType = "error"
	SendMessageTypeOffer          SendMessageType = "offer"
	SendMessageTypeAnswer         SendMessageType = "answer"
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
	log.Println("error:", msg)
	return &SendMessage{
		Type:    SendMessageTypeError,
		Payload: []byte(msg.Error()),
	}
}
