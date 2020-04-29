package repository

import (
	"errors"

	"github.com/tockn/singo/model"
)

type Room interface {
	Get(roomID string) (*model.Room, error)
	Update(r *model.Room) (*model.Room, error)
	Create(r *model.Room) (*model.Room, error)
	GetByClientID(clientID string) (*model.Room, error)
}

var (
	ErrNotFound = errors.New("not found")
)
