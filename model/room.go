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

// var Rooms = map[int]*Room{}
type Rooms []Room

var RoomToHub = map[int]*websocket.Hub{}

func (r *Room) Create(name string, user_id int) (*Room, error) {
	conn := db.DB.GetConnection()
	if err := conn.Create(r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Room) GetAll() (*Rooms, error) {
	conn := db.DB.GetConnection()
	var rooms Rooms
	if err := conn.Find(&rooms).Error; err != nil {
		return nil, err
	}
	return &rooms, nil
}
