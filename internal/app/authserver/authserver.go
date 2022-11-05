package authserver

import (
	"context"
	"sync"

	"github.com/VladimirBlinov/AuthService/internal/authservice"
	"github.com/VladimirBlinov/AuthService/internal/store"
)

type SessionManager struct {
	mu           sync.RWMutex
	sessionStore store.Store
	authservice.UnimplementedAuthServiceServer
}

func NewSessionManager(unimplementedAuthServiceServer authservice.UnimplementedAuthServiceServer, store store.Store) *SessionManager {
	return &SessionManager{
		mu:                             sync.RWMutex{},
		sessionStore:                   store,
		UnimplementedAuthServiceServer: unimplementedAuthServiceServer,
	}
}

func (sm *SessionManager) Create(ctx context.Context, s *authservice.Session) (*authservice.SessionID, error) {
	return sm.sessionStore.Session().Create(s)
}

func (sm *SessionManager) Check(ctx context.Context, sID *authservice.SessionID) (*authservice.Session, error) {
	return sm.sessionStore.Session().Check(sID)
}

func (sm *SessionManager) Delete(ctx context.Context, sID *authservice.SessionID) (*authservice.Nothing, error) {
	return sm.sessionStore.Session().Delete(sID)
}
