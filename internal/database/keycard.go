package database

import (
	"gorm.io/gorm"
	"time"
)

type KeyCardModel struct {
	DB *gorm.DB
}
type KeyCard struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint `gorm:"index"`
	PosDeviceID uint `gorm:"index"`
	CreatedAt   time.Time
	ExpiredAt   *time.Time
}
