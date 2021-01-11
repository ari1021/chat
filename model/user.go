package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           string `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Name         string         `form:"name" validate:"required,excludesall= " gorm:"unique;not null"`
	idToken      string
	accessToken  string
	refreshToken string
}

var Users = map[string]*User{}
