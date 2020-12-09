package model

type group struct {
	id      int
	name    string
	members []*User
}

var groups = map[int]*group{}
