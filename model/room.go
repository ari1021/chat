package model

import (
	"github.com/ari1021/websocket/server/websocket"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model        // equal ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `form:"name" validate:"required,excludesall= " gorm:"uniqueIndex"`
	UserID     string
	User       User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Members    []*User
}

var Rooms = map[int]*Room{}

var RoomToHub = map[int]*websocket.Hub{}
