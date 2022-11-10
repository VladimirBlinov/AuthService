package store

import (
	"github.com/VladimirBlinov/AuthService/internal/authservice"
	"github.com/VladimirBlinov/AuthService/internal/model"
)

type SessionRepo interface {
	Create(*model.Session) (*model.SessionID, error)
	Check(*model.SessionID) (*model.Session, error)
	Delete(*authservice.SessionID) (*authservice.Nothing, error)
}
