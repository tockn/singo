package handler

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/tockn/singo/model"
)

// HandleMessageはWebsocketで来たメッセージのTypeを元に適切なHandle関数を実行する
func (h *Handler) HandleReceiveMessage(c *model.Client, conn *websocket.Conn) {
	defer func() {
		_ = conn.Close()
	}()
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		// Unmarshalして適切なmethod実行
		var req ReceiveMessage
		if err := json.Unmarshal(msg, &req); err != nil {
			// send error message
			continue
		}

		var resp *SendMessage
		switch req.Type {
		case ReceiveMessageTypeJoinRoom:
			resp = h.handleJoinRoom(c, req.Payload)
		case ReceiveMessageTypeOffer:
			resp = h.handleSDPOffer(c, req.Payload)
		case ReceiveMessageTypeAnswer:
			resp = h.handleSDPAnswer(c, req.Payload)
		default:
			// send bad request
			// invalid type
			resp = newErrorMessage(ErrMsgInvalidType)
		}

		if resp == nil {
			continue
		}
		respMsg, err := json.Marshal(resp)
		if err != nil {
			continue
		}
		if err := sendMessage(conn, respMsg); err != nil {
			continue
		}
	}
}

type MessageJoinRoom struct {
	RoomID string `json:"room_id"`
}

func (h *Handler) handleJoinRoom(c *model.Client, msg []byte) *SendMessage {
	var req MessageJoinRoom
	if err := json.Unmarshal(msg, &req); err != nil {
		return newErrorMessage(ErrMsgInvalidPayload)
	}
	if err := h.manager.JoinRoom(c, req.RoomID); err != nil {
		return newErrorMessage(ErrMsgInvalidPayload)
	}
	return nil
}

type MessageSDPOffer struct {
	SDP      *model.SDP `json:"sdp"`
	ClientID string     `json:"client_id"`
}

// SDP offerが来た時に呼ばれる。Room IDの
func (h *Handler) handleSDPOffer(c *model.Client, msg []byte) *SendMessage {
	var req MessageSDPOffer
	if err := json.Unmarshal(msg, &req); err != nil {
		return newErrorMessage(ErrMsgInvalidPayload)
	}
	if err := h.manager.TransferSDPOffer(c, req.SDP, req.ClientID); err != nil {
		return newErrorMessage(ErrMsgInternalError)
	}
	return nil
}

type MessageSDPAnswer struct {
	SDP      *model.SDP `json:"sdp"`
	ClientID string     `json:"client_id"`
}

func (h *Handler) handleSDPAnswer(c *model.Client, msg []byte) *SendMessage {
	var req MessageSDPAnswer
	if err := json.Unmarshal(msg, &req); err != nil {
		return newErrorMessage(ErrMsgInvalidPayload)
	}
	if err := h.manager.TransferSDPAnswer(c, req.SDP, req.ClientID); err != nil {
		return newErrorMessage(ErrMsgInternalError)
	}
	return nil
}
