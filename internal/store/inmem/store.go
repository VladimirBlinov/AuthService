package inmem

import (
	"github.com/VladimirBlinov/AuthService/internal/authservice"
	"github.com/VladimirBlinov/AuthService/internal/store"
)

// Store
type Store struct {
	sessionRepo *SessionRepo
}

// Store constructor
func New() *Store {
	return &Store{}
}

func (s *Store) Session() store.SessionRepo {
	if s.sessionRepo != nil {
		return s.sessionRepo
	}

	s.sessionRepo = &SessionRepo{
		store:    s,
		sessions: make(map[*authservice.SessionID]authservice.Session),
	}
	return s.sessionRepo
}
