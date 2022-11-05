package authserver

import ("sync")

type SessionManager struct {
	mu           sync.RWMutex
	sessionStore store.Session
}