package model

import "github.com/rs/xid"

type Room struct {
	ID      string
	Name    string
	Clients map[string]*Client
}

func NewRoom(name string) *Room {
	return &Room{
		ID:      name,
		Name:    name,
		Clients: make(map[string]*Client, 0),
	}
}

type Client struct {
	ID   string
	Name string
	Send chan *Message
}

func NewClient(name string) *Client {
	return &Client{
		ID:   xid.New().String(),
		Name: name,
		Send: make(chan *Message, 16),
	}
}
