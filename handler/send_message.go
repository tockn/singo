package handler

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/tockn/singo/model"
)

func (h *Handler) HandleSendMessage(ctx context.Context, c *model.Client, conn *websocket.Conn) {
	t := time.NewTicker(1 * time.Second)
	defer func() {
		t.Stop()
		_ = conn.Close()
		_ = h.manager.LeaveRoom(c)
	}()
	for {
		select {
		case msg := <-c.Send:
			bs, err := json.Marshal(msg)
			if err != nil {
				log.Println(err)
				return
			}
			if err := sendMessage(conn, bs); err != nil {
				log.Println(err)
				return
			}
		case <-t.C:
			if err := sendMessage(conn, []byte(`{"type":"ping"}`)); err != nil {
				log.Println("disconnect: ", c.ID)
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
