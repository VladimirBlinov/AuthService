package store

import (
	"github.com/VladimirBlinov/AuthService/internal/model"
)

type SessionRepo interface {
	Create(*model.Session) (*model.SessionID, error)
	Check(*model.SessionID) (*model.Session, error)
	Delete(*model.SessionID) (bool, error)
}
