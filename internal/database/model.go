package database

import "gorm.io/gorm"

type Models struct {
	Users     *UserModel
	Terminals *TerminalModel
	KeyCards  *KeyCardModel
}

func NewModels(db *gorm.DB) *Models {
	return &Models{
		Users:     &UserModel{db},
		Terminals: &TerminalModel{db},
		KeyCards:  &KeyCardModel{db},
	}
}
