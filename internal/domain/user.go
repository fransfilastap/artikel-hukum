package domain

import (
	"gorm.io/gorm"
	"time"
)

type Role string
type UserID uint

var (
	Admin  Role = "admin"
	Editor Role = "editor"
	Author Role = "author"
	NoRole Role = "no-role"
)

type User struct {
	Id              uint
	FullName        string    `gorm:"type:varchar(255);not null"`
	Password        string    `gorm:"not null"`
	Email           string    `gorm:"unique;not null"`
	EmailVerifiedAt time.Time `gorm:"nullable;default:null"`
	Avatar          string    `gorm:"nullable;default:null"`
	Role            Role      `gorm:"type:varchar(10);enum('admin','editor', 'author', 'no-role');default:'no-role';not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
