package model

type SDP string

type Room struct {
	ID      string
	Name    string
	Clients map[string]*Client
}

type Client struct {
	ID   string
	Name string
	Send chan []byte
}
