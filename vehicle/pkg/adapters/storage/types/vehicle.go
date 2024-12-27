package types

import (
	"time"

	"gorm.io/gorm"
)

type Vehicle struct {
	Id              string `gorm:"primaryKey"`
	Unicode         string
	RequiredExperts int32
	Speed           int32
	RentPrice       int32
	IsActive        bool
	Type            string
	OwnerID         uint64
	Passenger       int
	Model           int
	Owner           *Owner         `gorm:"foreignKey:OwnerID"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"` // Auto-set when created
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"` // Auto-set when updated
	DeletedAt       gorm.DeletedAt `gorm:"index"`          // Soft delete (optional)
}

type Owner struct {
	Id        uint64 `gorm:"primaryKey"` // Primary key
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time      `gorm:"autoCreateTime"` // Auto-set when created
	UpdatedAt time.Time      `gorm:"autoUpdateTime"` // Auto-set when updated
	DeletedAt gorm.DeletedAt `gorm:"index"`          // Soft delete (optional)
}
