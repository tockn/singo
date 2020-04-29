package mem

import (
	"sync"

	"github.com/tockn/singo/model"
	"github.com/tockn/singo/repository"
)

func NewRoomRepository() repository.Room {
	return &roomRepository{
		rooms: make(map[string]*model.Room, 0),
	}
}

type roomRepository struct {
	mutex sync.Mutex
	rooms map[string]*model.Room
}

func (re *roomRepository) Get(roomID string) (*model.Room, error) {
	re.mutex.Lock()
	defer re.mutex.Unlock()
	r, ok := re.rooms[roomID]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return r, nil
}

func (re *roomRepository) Update(r *model.Room) (*model.Room, error) {
	re.mutex.Lock()
	defer re.mutex.Unlock()
	if _, ok := re.rooms[r.ID]; !ok {
		return nil, repository.ErrNotFound
	}
	re.rooms[r.ID] = r
	return r, nil
}

func (re *roomRepository) Create(r *model.Room) (*model.Room, error) {
	re.mutex.Lock()
	defer re.mutex.Unlock()
	re.rooms[r.ID] = r
	return r, nil
}

func (re *roomRepository) GetByClientID(clientID string) (*model.Room, error) {
	re.mutex.Lock()
	defer re.mutex.Unlock()
	for _, r := range re.rooms {
		for _, c := range r.Clients {
			if c.ID != clientID {
				continue
			}
			return r, nil
		}
	}
	return nil, repository.ErrNotFound
}
