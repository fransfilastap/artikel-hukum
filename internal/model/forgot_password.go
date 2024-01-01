package model

import "time"

type PasswordResetToken struct {
	UserID uint   `gorm:"user_id;primaryKey;autoIncrement:false"`
	Token  string `gorm:"primaryKey;type:varchar(128);unique"`
	Expiry time.Time
}
