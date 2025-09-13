package database

import (
	"gorm.io/gorm"
)

type UserModel struct {
	DB *gorm.DB
}
type User struct {
	ID       uint   `gorm:"primaryKey"`
	FullName string `gorm:"size:128"`
	//Status    string `gorm:"size:16;default:ACTIVE"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
	Devices  []Terminal `gorm:"foreignKey:UserID"`
	CardKeys []KeyCard  `gorm:"foreignKey:UserID"`
}
