package inmem

import (
	"sync"

	"github.com/VladimirBlinov/AuthService/internal/model"
	"github.com/VladimirBlinov/AuthService/internal/store"
)

// Store
type Store struct {
	mu          sync.RWMutex
	sessionRepo *SessionRepo
}

// Store constructor
func New() *Store {
	return &Store{
		mu: sync.RWMutex{},
	}
}

func (s *Store) Session() store.SessionRepo {
	if s.sessionRepo != nil {
		return s.sessionRepo
	}

	s.sessionRepo = &SessionRepo{
		store:    s,
		sessions: make(map[model.SessionID]*model.Session),
	}
	return s.sessionRepo
}
