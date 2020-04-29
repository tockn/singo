package manager

import (
	"github.com/tockn/singo/model"
	"github.com/tockn/singo/repository"
)

type Manager struct {
	roomRepo repository.Room
}

func (m *Manager) CreateRoom(name string) (*model.Room, error) {
	r := model.NewRoom(name)
	return m.roomRepo.Create(r)
}

func (m *Manager) JoinRoom(c *model.Client, roomID string) error {
	r, err := m.roomRepo.Get(roomID)
	if err != nil {
		return err
	}
	r.Clients[c.ID] = c
	if _, err := m.roomRepo.Update(r); err != nil {
		return err
	}
	return nil
}

type SDPOfferPayload struct {
	SDP *model.SDP
}

func (m *Manager) TransferSDPOffer(senderClient *model.Client, sdp *model.SDP) error {
	r, err := m.roomRepo.GetByClientID(senderClient.ID)
	if err != nil {
		return err
	}
	msg := &model.Message{
		Type:    model.MessageTypeSDPOffer,
		Payload: &SDPOfferPayload{SDP: sdp},
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
	SDP *model.SDP
}

func (m *Manager) TransferSDPAnswer(senderClient *model.Client, sdp *model.SDP) error {
	r, err := m.roomRepo.GetByClientID(senderClient.ID)
	if err != nil {
		return err
	}
	msg := &model.Message{
		Type:    model.MessageTypeSDPAnswer,
		Payload: &SDPAnswerPayload{SDP: sdp},
	}
	for _, c := range r.Clients {
		if c.ID == senderClient.ID {
			continue
		}
		c.Send <- msg
	}
	return nil
}
