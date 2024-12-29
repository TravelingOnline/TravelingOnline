package entities

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique;not null"`
	Type      string `gorm:"not null"` // E.g., "read-only", "write-only"
	Scope     string `gorm:"not null"` // E.g., "user", "admin"
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
