package manager

import "github.com/tockn/singo/model"

type Manager struct{}

func (m *Manager) CreateRoom(name string) (*model.Room, error) {
	return nil, nil
}

func (m *Manager) JoinRoom(clientName string) error {
	return nil
}

func (m *Manager) TransferSDPOffer(senderClient *model.Client, sdp *model.SDP) error {
	return nil
}

func (m *Manager) TransferSDPAnswer(senderClient *model.Client, sdp *model.SDP) error {
	return nil
}
