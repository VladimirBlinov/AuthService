package authserver

import (
	"context"
	"sync"

	"github.com/VladimirBlinov/AuthService/internal/authservice"
	"github.com/VladimirBlinov/AuthService/internal/model"
	"github.com/VladimirBlinov/AuthService/internal/store"
	"github.com/sirupsen/logrus"
)

type SessionManager struct {
	mu           sync.RWMutex
	sessionStore store.Store
	logger       *logrus.Logger
	authservice.UnimplementedAuthServiceServer
}

func NewSessionManager(unimplementedAuthServiceServer authservice.UnimplementedAuthServiceServer, store store.Store, logger *logrus.Logger) *SessionManager {
	return &SessionManager{
		mu:                             sync.RWMutex{},
		sessionStore:                   store,
		logger:                         logger,
		UnimplementedAuthServiceServer: unimplementedAuthServiceServer,
	}
}

func (sm *SessionManager) Create(ctx context.Context, s *authservice.Session) (*authservice.SessionID, error) {
	session := model.NewSession(s)
	sessionID, err := sm.sessionStore.Session().Create(session)
	if err != nil {
		sm.logger.Errorf("SM create error: %s", err)
		return nil, err
	}

	sm.logger.Infof("SM created sessionId: %s", sessionID.SessionID)
	return &authservice.SessionID{
		ID: sessionID.SessionID,
	}, nil
}

func (sm *SessionManager) Check(ctx context.Context, sID *authservice.SessionID) (*authservice.Session, error) {
	sessionID := model.NewSessionID(sID)

	session, err := sm.sessionStore.Session().Check(sessionID)
	if err != nil {
		sm.logger.Errorf("SM check error: %s", err)
		return nil, err
	}

	sm.logger.Infof("SM checked session UserID: %d", session.UserID)
	return &authservice.Session{
		UserID: int32(session.UserID),
	}, nil
}

func (sm *SessionManager) Delete(ctx context.Context, sID *authservice.SessionID) (*authservice.Nothing, error) {
	return sm.sessionStore.Session().Delete(sID)
}
