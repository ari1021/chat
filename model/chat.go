package model

import (
	"time"

	"gorm.io/gorm"
)

type IChat interface {
	Find(roomID int, limit int, offset int) (*Chats, error)
	Create(message string, roomID int, userName string) (*Chat, error)
}
type ChatRepository struct {
	conn *gorm.DB
}
type Chat struct {
	ID        int       `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"index"`
	RoomID    int
	Room      Room   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Message   string `gorm:"not null"`
	UserName  string `gorm:"not null"`
}

type Chats []Chat

func (cr ChatRepository) Find(roomID int, limit int, offset int) (*Chats, error) {
	cs := &Chats{}
	if err := cr.conn.Order("created_at desc").Limit(limit).Offset(offset).Find(cs, "room_id = ?", roomID).Error; err != nil {
		return nil, err
	}
	return cs, nil
}

func (cr ChatRepository) Create(message string, roomID int, userName string) (*Chat, error) {
	c := &Chat{
		RoomID:   roomID,
		Message:  message,
		UserName: userName,
	}
	if err := cr.conn.Create(c).Error; err != nil {
		return nil, err
	}
	return c, nil
}
