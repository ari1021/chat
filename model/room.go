package model

import "github.com/ari1021/websocket/server/websocket"

type Room struct {
	ID      int
	Name    string `form:"name" validate:"required,excludesall= "`
	Members []*User
}

var Rooms = map[int]*Room{}

var RoomToHub = map[int]*websocket.Hub{}
