package entities

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID         uint    `gorm:"primaryKey"`
	Name       string  `gorm:"unique;not null"`
	Users      []*User `gorm:"many2many:user_roles;"` // Many-to-Many relationship
	Permisions []*Permission
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
