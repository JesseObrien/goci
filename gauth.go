package gauth

import ()

type Hasher interface {
	make(string password) []byte
}

type Hash struct{}

func (h *Hash) make(password string) []byte {
	return []byte{}
}

type User struct {
	identifier interface{}
	password   interface{}
}

type Authenticator interface {
	Login(credentials []interface{}) bool
	Logout() bool
	Check() bool
	Attempt(credentials []interface{}) bool
	User() interface{}
	Validate(credentials []interface{}) bool
	Once(credentials []interface{}) bool
}

type Gauth struct {
}
