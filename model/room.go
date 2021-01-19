package model

import (
	"github.com/ari1021/websocket/db"
	"github.com/ari1021/websocket/server/websocket"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model        // equal ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `form:"name" validate:"required,excludesall= " gorm:"size:255;uniqueIndex"`
	UserID     int
	User       User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Members    []*User
}

var Rooms = map[int]*Room{}

var RoomToHub = map[int]*websocket.Hub{}

func (r *Room) Create(name string, user_id int) (*Room, error) {
	conn := db.DB.GetConnection()
	res := conn.Create(r)
	if err := res.Error; err != nil {
		return nil, err
	}
	return r, nil
}
