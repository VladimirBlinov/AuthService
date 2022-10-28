package store

type Store interface{
	Session() SessionRepo
}