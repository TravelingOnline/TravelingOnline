package types

import (
	"time"

	"gorm.io/gorm"
)

type Tour struct {
	Id            string `gorm:"primaryKey;type:uuid;"`
	Source        string `gorm:"index"`
	Destination   string `gorm:"index"`
	StartDate     string `gorm:"not null"`
	EndDate       string `gorm:"not null"`
	Type          string
	Price         int32
	Capacity      int
	VehicleID     string  `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Cascade updates, set null on delete
	CompanyID     string  `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`  // Cascade updates and deletes
	Company       Company `gorm:"foreignKey:CompanyID"`
	AdminApprove  bool
	Ended         bool
	Vehicle       *Vehicle         `gorm:"foreignKey:VehicleID"`
	TechnicalTeam []*TechnicalTeam `gorm:"many2many:tour_technical_teams;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Cascade updates and deletes for many-to-many
	CreatedAt     time.Time        `gorm:"autoCreateTime"`
	UpdatedAt     time.Time        `gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt   `gorm:"index"`
}

type TechnicalTeam struct {
	Id        string `gorm:"primaryKey"`
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
