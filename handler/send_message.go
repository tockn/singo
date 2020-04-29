package handler

import (
	"github.com/gorilla/websocket"
	"github.com/tockn/singo/model"
)

func (h *Handler) HandleSendMessage(c *model.Client, conn *websocket.Conn) {
	defer func() {
		_ = conn.Close()
	}()
	for {
		select {
		case msg := <-c.Send:
			if err := sendMessage(conn, msg); err != nil {
				return
			}
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
