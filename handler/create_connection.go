package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/tockn/singo/model"
)

type SendClientID struct {
	ClientID string `json:"client_id"`
}

// CreateConnectionでwebsocketコネクションを確立
// Clientを作成してレスポンスとして通知する
func (h *Handler) CreateConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	c := model.NewClient("name")
	ctx := context.Background()
	go h.HandleSendMessage(ctx, c, conn)
	go h.HandleReceiveMessage(c, conn)
	resp := &SendClientID{ClientID: c.ID}
	payload, _ := json.Marshal(resp)
	body, _ := json.Marshal(SendMessage{
		Type:    SendMessageTypeNotifyClientID,
		Payload: payload,
	})
	sendMessage(conn, body)
}
