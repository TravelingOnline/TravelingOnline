package types

import (
	"time"

	"gorm.io/gorm"
)

type Tour struct {
	Id              string `gorm:"primaryKey"`
	Source          string
	Destination     string
	StartDate       string
	EndDate         string
	Type            string
	Price           int32
	VehicleID       string
	CompanyID       string  `gorm:"not null"`
	Company         Company `gorm:"foreignKey:CompanyID;references:Id"`
	AdminApprove    bool
	Ended           bool
	Vehicle         *Vehicle `gorm:"foreignKey:VehicleID"`
	TechnicalTeamID string
	TechnicalTeam   []*TechnicalTeam `gorm:"foreignKey:TechnicalTeamID"`
	CreatedAt       time.Time        `gorm:"autoCreateTime"` // Auto-set when created
	UpdatedAt       time.Time        `gorm:"autoUpdateTime"` // Auto-set when updated
	DeletedAt       gorm.DeletedAt   `gorm:"index"`          // Soft delete (optional)
}

type TechnicalTeam struct {
	Id        uint64 `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Age       int32
	Expertise string
	CreatedAt time.Time      `gorm:"autoCreateTime"` // Auto-set when created
	UpdatedAt time.Time      `gorm:"autoUpdateTime"` // Auto-set when updated
	DeletedAt gorm.DeletedAt `gorm:"index"`          // Soft delete (optional)
}

type Vehicle struct {
	Id              string `gorm:"primaryKey"`
	Unicode         string
	RequiredExperts int32
	Speed           int32
	RentPrice       int32
	Type            string
	Passenger       int32
	Model           int32
	CreatedAt       time.Time      `gorm:"autoCreateTime"` // Auto-set when created
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"` // Auto-set when updated
	DeletedAt       gorm.DeletedAt `gorm:"index"`          // Soft delete (optional)
}
