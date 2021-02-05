package model

import (
	"github.com/ari1021/websocket/server/websocket"
	"gorm.io/gorm"
)

type IRoom interface {
	Create(name string) (*Room, error)
	FindAll() (*Rooms, error)
	Delete(id uint) (*Room, error)
}

type Room struct {
	gorm.Model        // equal ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `form:"name" validate:"required,excludesall= " gorm:"size:255;uniqueIndex"`
}

type RoomRepository struct {
	conn *gorm.DB
}

func NewRoomRepository(conn *gorm.DB) *RoomRepository {
	return &RoomRepository{conn: conn}
}

// var Rooms = map[int]*Room{}
type Rooms []Room

var RoomToHub = map[uint]*websocket.Hub{}

func (rr RoomRepository) Create(name string) (*Room, error) {
	r := &Room{Name: name}
	if err := rr.conn.Create(r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func (rr RoomRepository) FindAll() (*Rooms, error) {
	rs := &Rooms{}
	if err := rr.conn.Find(rs).Error; err != nil {
		return nil, err
	}
	return rs, nil
}

func (rr RoomRepository) Delete(id uint) (*Room, error) {
	r := &Room{
		Model: gorm.Model{
			ID: id,
		},
	}
	result := rr.conn.Delete(r)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	} else if err := result.Error; err != nil {
		return nil, err
	}
	return r, nil
}
