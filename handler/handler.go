package handler

import (
	"encoding/json"

	"github.com/tockn/singo/manager"

	"github.com/gorilla/websocket"
)

type Handler struct {
	manager *manager.Manager
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type RequestType string

var (
	RequestTypeOffer  RequestType = "offer"
	RequestTypeAnswer RequestType = "answer"
)

type RequestMessage struct {
	Type    RequestType     `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type ResponseType string

var (
	ResponseTypeError  ResponseType = "error"
	ResponseTypeOffer  ResponseType = "offer"
	ResponseTypeAnswer ResponseType = "answer"
)

type ResponseMessage struct {
	Type    ResponseType    `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

func newErrorResponse(payload json.RawMessage) *ResponseMessage {
	return &ResponseMessage{
		Type:    ResponseTypeError,
		Payload: payload,
	}
}
