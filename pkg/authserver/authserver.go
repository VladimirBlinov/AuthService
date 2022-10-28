package authserver

type SessionManager struct {
	mu sync.RWMutex
	sessionStore sessionStore
}