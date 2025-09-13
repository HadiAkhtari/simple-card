package database

import (
	"gorm.io/gorm"
	"time"
)

type TerminalModel struct {
	DB *gorm.DB
}
type Terminal struct {
	ID           uint   `gorm:"primaryKey"`
	UserID       uint   `gorm:"index"`
	SerialNumber string `gorm:"uniqueIndex;size:64;not null"`
	Model        string `gorm:"size:64"`
	Status       string `gorm:"size:20;default:ACTIVE"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CardKeys     []KeyCard `gorm:"foreignKey:PosDeviceID"`
}
