package store

type SessionRepo interface{
	Create(Session)(SessionID)
	Check(SessionID)(Session)
	Delete(SessionID)(Nothing)
}