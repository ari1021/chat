package model

type Room struct {
	ID      int
	Name    string
	Members []*User
}

var Rooms = map[int]*Room{}
