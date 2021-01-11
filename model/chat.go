package model

import (
	"time"
)

type Chat struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	RoomID    int
	Room      Room `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    int
	User      User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Message   string `gorm:not null`
}
