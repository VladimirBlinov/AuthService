package inmem

import (
	"errors"
	"fmt"

	"github.com/VladimirBlinov/AuthService/internal/model"
	"github.com/google/uuid"
)

type SessionRepo struct {
	store    *Store
	sessions map[model.SessionID]*model.Session
}

func (sr *SessionRepo) Create(as *model.Session) (*model.SessionID, error) {
	sessionID := &model.SessionID{
		SessionID: uuid.New().String(),
	}
	sr.store.mu.Lock()
	defer sr.store.mu.Unlock()
	sr.sessions[*sessionID] = as

	return sessionID, nil
}

func (sr *SessionRepo) Check(asID *model.SessionID) (*model.Session, error) {
	sr.store.mu.RLock()
	defer sr.store.mu.RUnlock()
	as, ok := sr.sessions[*asID]
	if !ok {
		return nil, errors.New(fmt.Sprintf("session not found by sessionID %s", asID.SessionID))
	}
	return as, nil
}

func (sr *SessionRepo) Delete(asID *model.SessionID) (bool, error) {
	sr.store.mu.Lock()
	defer sr.store.mu.Unlock()

	_, deleted := sr.sessions[*asID]
	if !deleted {
		return false, errors.New(fmt.Sprintf("session not found by sessionID %s", asID.SessionID))
	}
	delete(sr.sessions, *asID)
	return deleted, nil
}
