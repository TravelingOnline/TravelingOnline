package types

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `gorm:"column:email;unique;not null"`
	PasswordHash string `gorm:"column:password_hash;not null"`
}
