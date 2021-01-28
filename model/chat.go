package model

import (
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	ID        int       `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"index"`
	RoomID    int
	Room      Room   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Message   string `gorm:"not null"`
}

type Chats []Chat

func (c *Chats) Find(conn *gorm.DB, roomID int, limit int, offset int) (*Chats, error) {
	if err := conn.Order("created_at desc").Limit(limit).Offset(offset).Find(c, "room_id = ?", roomID).Error; err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Chat) Create(conn *gorm.DB) (*Chat, error) {
	if err := conn.Create(c).Error; err != nil {
		return nil, err
	}
	return c, nil
}
