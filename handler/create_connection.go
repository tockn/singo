package handler

import (
	"net/http"

	"github.com/tockn/singo/model"
)

// CreateConnectionでwebsocketコネクションを確立
// Clientを作成してレスポンスとして通知する
func (h *Handler) CreateConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	c := &model.Client{}
	go h.HandleSendMessage(c, conn)
}
