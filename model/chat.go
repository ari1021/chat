package model

import (
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	RoomID    int
	Room      Room `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    int
	User      User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Message   string `gorm:"not null"`
}

type Chats []Chat

func (c *Chats) Find(conn *gorm.DB, groupID int, limit int, offset int) (*Chats, error) {
	if err := conn.Order("created_at　desc").Limit(limit).Offset(offset).Find(c, "room_id = ?", groupID).Error; err != nil {
		return nil, err
	}
	return c, nil
}
