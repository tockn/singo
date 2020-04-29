package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
)

type SDP string

type Room struct {
	ID      string
	Name    string
	Clients map[string]*Client
}

func CreateRoom(name string) (*Room, error) {
	return nil, nil
}

func Join(name string) error {
	return nil
}

func TransferSDPOffer(sdp *SDP, senderClientID string) error {
	return nil
}

func TransferSDPAnswer(sdp *SDP, senderClientID string) error {
	return nil
}

type Client struct {
	ID   string
	Name string
	Send chan []byte
}

// ListenMessage should be executed in goroutine
func (c *Client) ListenMessage() {
}

type Handler struct{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// CreateConnectionでwebsocketコネクションを確立
// Clientを作成してレスポンスとして通知する
func (h *Handler) CreateConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	c := &Client{}
	go h.HandleSendMessage(c, conn)
}

func (h *Handler) HandleSendMessage(c *Client, conn *websocket.Conn) {
	defer func() {
		_ = conn.Close()
	}()
	for {
		select {
		case msg := <-c.Send:
			w, err := conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(msg)
			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

// HandleMessageはWebsocketで来たメッセージのTypeを元に適切なHandle関数を実行する
func (h *Handler) HandleReceiveMessage(conn *websocket.Conn) {
	// Typeを元に、SDPOffer, SDPAnswerを実行する
	// 返り値がnilでなければ、そのResponseMessageを送信する
	defer func() {
		_ = conn.Close()
	}()
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		// Unmarshalして適切なmethod実行
	}
}

type RequestSDPOffer struct {
	RoomID string `json:"room_id"`
	SDP    SDP    `json:"sdp"`
}

// SDP offerが来た時に呼ばれる。Room IDの
func (h *Handler) SDPOffer(msg *RequestMessage) *ResponseMessage {
	return nil
}

func (h *Handler) SDPAnswer(msg *RequestMessage) *ResponseMessage {
	return nil
}

type RequestMessage struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}
type ResponseMessage struct{}

func main() {

}
