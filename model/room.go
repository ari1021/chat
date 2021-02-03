package model

import (
	"github.com/ari1021/websocket/server/websocket"
	"gorm.io/gorm"
)

type IRoom interface {
	Create(conn *gorm.DB) (*Room, error)
	FindAll(conn *gorm.DB) (*Rooms, error)
	Delete(conn *gorm.DB) (*Room, error)
}

type Room struct {
	gorm.Model        // equal ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `form:"name" validate:"required,excludesall= " gorm:"size:255;uniqueIndex"`
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

func (r *Rooms) FindAll(conn *gorm.DB) (*Rooms, error) {
	if err := conn.Find(r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Room) Delete(conn *gorm.DB) (*Room, error) {
	result := conn.Delete(r)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	} else if err := result.Error; err != nil {
		return nil, err
	}
	return r, nil
}
