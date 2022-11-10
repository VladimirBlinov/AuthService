package model

import (
	"testing"

	"github.com/VladimirBlinov/AuthService/internal/authservice"
)

func TestInputSession(t *testing.T) *authservice.Session {
	return &authservice.Session{
		UserID: 1,
	}
}

func TestSession(t *testing.T) *Session {
	return &Session{
		UserID: 1,
	}
}
