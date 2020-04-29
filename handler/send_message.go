package handler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/tockn/singo/manager"

	"github.com/gorilla/websocket"
	"github.com/tockn/singo/model"
)

func (h *Handler) HandleSendMessage(ctx context.Context, c *model.Client, conn *websocket.Conn) {
	defer func() {
		_ = conn.Close()
	}()
	for {
		select {
		case msg := <-c.Send:
			var err error
			switch msg.Type {
			case model.MessageTypeSDPOffer:
				err = sendSDPOffer(conn, msg)
			case model.MessageTypeSDPAnswer:
				err = sendSDPAnswer(conn, msg)
			}
			if err != nil {
				log.Println(err)
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

func sendMessage(conn *websocket.Conn, msg []byte) error {
	w, err := conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return err
	}
	w.Write(msg)
	if err := w.Close(); err != nil {
		return err
	}
	return nil
}

type SendMessageSDPOfferPayload struct {
	SDP *model.SDP `json:"sdp"`
}

func sendSDPOffer(conn *websocket.Conn, msg *model.Message) error {
	mp, ok := msg.Payload.(manager.SDPOfferPayload)
	if !ok {
		return ErrMsgInvalidPayload
	}
	p, err := json.Marshal(&SendMessageSDPOfferPayload{SDP: mp.SDP})
	if err != nil {
		return ErrMsgInvalidPayload
	}
	resp := &SendMessage{
		Type:    SendMessageTypeOffer,
		Payload: p,
	}
	respBody, err := json.Marshal(resp)
	if err != nil {
		return ErrMsgInvalidPayload
	}
	return sendMessage(conn, respBody)
}

type SendMessageSDPAnswerPayload struct {
	SDP *model.SDP `json:"sdp"`
}

func sendSDPAnswer(conn *websocket.Conn, msg *model.Message) error {
	mp, ok := msg.Payload.(manager.SDPAnswerPayload)
	if !ok {
		return ErrMsgInvalidPayload
	}
	p, err := json.Marshal(&SendMessageSDPAnswerPayload{SDP: mp.SDP})
	if err != nil {
		return ErrMsgInvalidPayload
	}
	resp := &SendMessage{
		Type:    SendMessageTypeAnswer,
		Payload: p,
	}
	respBody, err := json.Marshal(resp)
	if err != nil {
		return ErrMsgInvalidPayload
	}
	return sendMessage(conn, respBody)
}
