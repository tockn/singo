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
		var req RequestMessage
		if err := json.Unmarshal(msg, &req); err != nil {
			// send error message
			continue
		}

		var resp *ResponseMessage
		switch req.Type {
		case RequestTypeOffer:
			resp = h.sdpOffer(c, req.Payload)
		case RequestTypeAnswer:
			resp = h.sdpAnswer(c, req.Payload)
		default:
			// send bad request
			// invalid type
			resp = newErrorResponse(ErrMsgInvalidRequestType)
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

type RequestSDPOffer struct {
	SDP *model.SDP `json:"sdp"`
}

// SDP offerが来た時に呼ばれる。Room IDの
func (h *Handler) sdpOffer(c *model.Client, msg []byte) *ResponseMessage {
	var req RequestSDPOffer
	if err := json.Unmarshal(msg, &req); err != nil {
		return newErrorResponse(ErrMsgInvalidRequestPayload)
	}
	if err := h.manager.TransferSDPOffer(c, req.SDP); err != nil {
		return newErrorResponse(ErrMsgInternalError)
	}
	return nil

}

type RequestSDPAnswer struct {
	SDP *model.SDP `json:"sdp"`
}

func (h *Handler) sdpAnswer(c *model.Client, msg []byte) *ResponseMessage {
	var req RequestSDPAnswer
	if err := json.Unmarshal(msg, &req); err != nil {
		return newErrorResponse(ErrMsgInvalidRequestPayload)
	}
	if err := h.manager.TransferSDPAnswer(c, req.SDP); err != nil {
		return newErrorResponse(ErrMsgInternalError)
	}
	return nil
}
