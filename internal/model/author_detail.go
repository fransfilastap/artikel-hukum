package model

import (
	"gorm.io/gorm"
	"time"
)

type AuthorDetail struct {
	Id         uint
	UserID     uint `gorm:"user_id;not null"`
	User       User
	Occupation string         `gorm:"not null"`
	Company    string         `gorm:"not null"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
