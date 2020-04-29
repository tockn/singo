package main

import "net/http"

type SDP string

type Room struct {
	ID      string
	Name    string
	Clients []*Client
}

func CreateRoom(name string) (*Room, error) {
	return nil, nil
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
}

func (c *Client) JoinRoom(roomID string) (*Room, error) {
	return nil, nil
}

type Handler struct{}

func (h *Handler) CreateConnection(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) SDPOffer(msg *RequestMessage) *ResponseMessage {
	return nil
}

func (h *Handler) SDPAnswer(msg *RequestMessage) *ResponseMessage {
	return nil
}

type RequestMessage struct{}
type ResponseMessage struct{}

func main() {

}
