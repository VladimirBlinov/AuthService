package model

import "github.com/VladimirBlinov/AuthService/internal/authservice"

type Session struct {
	UserID int
}

func NewSession(s *authservice.Session) *Session {
	return &Session{
		UserID: int(s.UserID),
	}
}

type SessionID struct {
	SessionID string
}

func NewSessionID(s *authservice.SessionID) *SessionID {
	return &SessionID{
		SessionID: s.ID,
	}
}
