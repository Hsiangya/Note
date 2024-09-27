package models

type User struct {
	Id   string
	Name string
	Msg  chan string
}

func NewUser(name string, id string, msg chan string) *User {
	return &User{id, name, msg}
}
