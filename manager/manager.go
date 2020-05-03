package manager

import (
	"log"

	"github.com/tockn/singo/model"
	"github.com/tockn/singo/repository"
)

type Manager struct {
	roomRepo repository.Room
}

func NewManager(roomRepo repository.Room) *Manager {
	return &Manager{roomRepo: roomRepo}
}

func (m *Manager) CreateRoom(name string) (*model.Room, error) {
	r := model.NewRoom(name)
	return m.roomRepo.Create(r)
}

func (m *Manager) JoinRoom(c *model.Client, roomID string) error {
	r, err := m.roomRepo.Get(roomID)
	if err == repository.ErrNotFound {
		r = model.NewRoom(roomID)
		if _, err := m.roomRepo.Create(r); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	r.Clients[c.ID] = c
	if _, err := m.roomRepo.Update(r); err != nil {
		return err
	}
	log.Println("join: ", c.ID)
	return nil
}

type SDPOfferPayload struct {
	ClientID string
	SDP      *model.SDP
}

func (m *Manager) TransferSDPOffer(senderClient *model.Client, sdp *model.SDP) error {
	r, err := m.roomRepo.GetByClientID(senderClient.ID)
	if err != nil {
		return err
	}
	msg := &model.Message{
		Type:    model.MessageTypeSDPOffer,
		Payload: SDPOfferPayload{ClientID: senderClient.ID, SDP: sdp},
	}
	for _, c := range r.Clients {
		if c.ID == senderClient.ID {
			continue
		}
		c.Send <- msg
	}
	return nil
}

type SDPAnswerPayload struct {
	ClientID string
	SDP      *model.SDP
}

func (m *Manager) TransferSDPAnswer(senderClient *model.Client, sdp *model.SDP) error {
	r, err := m.roomRepo.GetByClientID(senderClient.ID)
	if err != nil {
		return err
	}
	msg := &model.Message{
		Type:    model.MessageTypeSDPAnswer,
		Payload: SDPAnswerPayload{ClientID: senderClient.ID, SDP: sdp},
	}
	for _, c := range r.Clients {
		if c.ID == senderClient.ID {
			continue
		}
		c.Send <- msg
	}
	return nil
}
