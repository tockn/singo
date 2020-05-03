package manager

import (
	"github.com/tockn/singo/model"
	"github.com/tockn/singo/repository"
)

type Room struct {
	roomRepo repository.Room
}

func NewManager(roomRepo repository.Room) *Room {
	return &Room{roomRepo: roomRepo}
}

func (rm *Room) CreateRoom(name string) (*model.Room, error) {
	r := model.NewRoom(name)
	return rm.roomRepo.Create(r)
}

func (rm *Room) JoinRoom(c *model.Client, roomID string) error {
	r, err := rm.roomRepo.Get(roomID)
	if err == repository.ErrNotFound {
		r = model.NewRoom(roomID)
		if _, err := rm.roomRepo.Create(r); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	r.Clients[c.ID] = c
	if _, err := rm.roomRepo.Update(r); err != nil {
		return err
	}
	return rm.notifyNewClient(roomID, c)
}

func (rm *Room) LeaveRoom(c *model.Client) error {
	r, err := rm.roomRepo.GetByClientID(c.ID)
	if err != nil {
		return err
	}
	go rm.notifyLeaveClient(r.ID, c)
	delete(r.Clients, c.ID)
	if _, err := rm.roomRepo.Update(r); err != nil {
		return err
	}
	return nil
}

type NewClientPayload struct {
	ClientID string `json:"client_id"`
}

func (rm *Room) notifyNewClient(roomID string, nc *model.Client) error {
	r, err := rm.roomRepo.Get(roomID)
	if err != nil {
		return err
	}
	msg := &model.Message{
		Type:    model.MessageTypeNewClient,
		Payload: NewClientPayload{ClientID: nc.ID},
	}
	for _, c := range r.Clients {
		if c.ID == nc.ID {
			continue
		}
		c.Send <- msg
	}
	return nil
}

type LeaveClientPayload struct {
	ClientID string `json:"client_id"`
}

func (rm *Room) notifyLeaveClient(roomID string, nc *model.Client) error {
	r, err := rm.roomRepo.Get(roomID)
	if err != nil {
		return err
	}
	msg := &model.Message{
		Type:    model.MessageTypeLeaveClient,
		Payload: LeaveClientPayload{ClientID: nc.ID},
	}
	for _, c := range r.Clients {
		if c.ID == nc.ID {
			continue
		}
		c.Send <- msg
	}
	return nil
}

type SDPOfferPayload struct {
	ClientID string     `json:"client_id"`
	SDP      *model.SDP `json:"sdp"`
}

func (rm *Room) TransferSDPOffer(senderClient *model.Client, sdp *model.SDP, clientID string) error {
	r, err := rm.roomRepo.GetByClientID(senderClient.ID)
	if err != nil {
		return err
	}
	msg := &model.Message{
		Type:    model.MessageTypeSDPOffer,
		Payload: SDPOfferPayload{ClientID: senderClient.ID, SDP: sdp},
	}
	for _, c := range r.Clients {
		if c.ID != clientID {
			continue
		}
		c.Send <- msg
	}
	return nil
}

type SDPAnswerPayload struct {
	ClientID string     `json:"client_id"`
	SDP      *model.SDP `json:"sdp"`
}

func (rm *Room) TransferSDPAnswer(senderClient *model.Client, sdp *model.SDP, clientID string) error {
	r, err := rm.roomRepo.GetByClientID(senderClient.ID)
	if err != nil {
		return err
	}
	msg := &model.Message{
		Type:    model.MessageTypeSDPAnswer,
		Payload: SDPAnswerPayload{ClientID: senderClient.ID, SDP: sdp},
	}
	for _, c := range r.Clients {
		if c.ID != clientID {
			continue
		}
		c.Send <- msg
	}
	return nil
}
