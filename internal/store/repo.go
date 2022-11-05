package store

import "github.com/VladimirBlinov/AuthService/internal/authservice"

type SessionRepo interface {
	Create(*authservice.Session) (*authservice.SessionID, error)
	Check(*authservice.SessionID) (*authservice.Session, error)
	Delete(*authservice.SessionID) (*authservice.Nothing, error)
}
