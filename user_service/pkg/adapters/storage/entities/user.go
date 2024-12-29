package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint    `gorm:"primaryKey"`
	FullName     string  `gorm:"not null"`
	Email        string  `gorm:"unique;not null"`
	Password     string  `gorm:"not null"`
	NationalCode string  `gorm:"unique;not null"`
	Roles        []*Role `gorm:"many2many:user_roles;"` // Many-to-Many relationship
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
