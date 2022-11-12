package authserver

import (
	"context"
	"fmt"

	"github.com/VladimirBlinov/AuthService/internal/model"
	"github.com/VladimirBlinov/AuthService/internal/store"
	"github.com/VladimirBlinov/AuthService/pkg/authservice"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type SessionManager struct {
	sessionStore store.Store
	logger       *logrus.Logger
	authservice.UnimplementedAuthServiceServer
}

func NewSessionManager(unimplementedAuthServiceServer authservice.UnimplementedAuthServiceServer, store store.Store, logger *logrus.Logger) *SessionManager {
	return &SessionManager{
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
		return nil, grpc.Errorf(codes.Internal, "SM create error: %s", err)
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
		return nil, grpc.Errorf(codes.NotFound, fmt.Sprintf("SM check error: %s", err))
	}

	sm.logger.Infof("SM SessionID %s: checked", sID.ID)
	return &authservice.Session{
		UserID: int32(session.UserID),
	}, nil
}

func (sm *SessionManager) Delete(ctx context.Context, sID *authservice.SessionID) (*authservice.Nothing, error) {
	sessionID := model.NewSessionID(sID)

	deleted, err := sm.sessionStore.Session().Delete(sessionID)
	if err != nil {
		sm.logger.Errorf("SM delete error: %s", err)
		return &authservice.Nothing{
			Dummy: deleted,
		}, grpc.Errorf(codes.NotFound, fmt.Sprintf("SM delete error: %s", err))
	}

	sm.logger.Infof("SM SessionID %s: deleted", sID.ID)
	return &authservice.Nothing{
		Dummy: deleted,
	}, nil
}
