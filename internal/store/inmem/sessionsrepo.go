package inmem

import "github.com/VladimirBlinov/AuthService/internal/authservice"

type SessionRepo struct {
	store    *Store
	sessions map[*authservice.SessionID]authservice.Session
}

func (sr *SessionRepo) Create(as *authservice.Session) (*authservice.SessionID, error) {

	return &authservice.SessionID{}, nil
}

func (sr *SessionRepo) Check(asID *authservice.SessionID) (*authservice.Session, error) {
	return &authservice.Session{}, nil
}

func (sr *SessionRepo) Delete(asID *authservice.SessionID) (*authservice.Nothing, error) {
	return &authservice.Nothing{}, nil
}
