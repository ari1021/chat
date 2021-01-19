package model

import (
	"github.com/ari1021/websocket/server/websocket"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model        // equal ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `form:"name" validate:"required,excludesall= " gorm:"size:255;uniqueIndex"`
	UserID     string
	User       User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Members    []*User
}

// var Rooms = map[int]*Room{}
type Rooms []Room

var RoomToHub = map[uint]*websocket.Hub{}

func (r *Room) Create(conn *gorm.DB) (*Room, error) {
	if err := conn.Create(r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Room) GetAll(conn *gorm.DB) (*Rooms, error) {
	var rooms Rooms
	if err := conn.Find(&rooms).Error; err != nil {
		return nil, err
	}
	return &rooms, nil
}

func (r *Room) Delete(conn *gorm.DB) (*Room, error) {
	if err := conn.Delete(r).Error; err != nil {
		return nil, err
	}
	return r, nil
}
