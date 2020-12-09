package model

type Room struct {
	ID      int
	Name    string `form:"name" validate:"required"`
	Members []*User
}

var Rooms = map[int]*Room{}
