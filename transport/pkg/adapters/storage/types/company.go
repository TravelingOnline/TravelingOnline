package types

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	Id        string `gorm:"primaryKey"`
	Name      string
	OwnerID   uint64
	Owner     *Owner         `gorm:"foreignKey:OwnerID"`
	CreatedAt time.Time      `gorm:"autoCreateTime"` // Auto-set when created
	UpdatedAt time.Time      `gorm:"autoUpdateTime"` // Auto-set when updated
	DeletedAt gorm.DeletedAt `gorm:"index"`          // Soft delete (optional)
}

type Owner struct {
	Id        string `gorm:"primaryKey"` // Primary key
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time      `gorm:"autoCreateTime"` // Auto-set when created
	UpdatedAt time.Time      `gorm:"autoUpdateTime"` // Auto-set when updated
	DeletedAt gorm.DeletedAt `gorm:"index"`          // Soft delete (optional)
}
