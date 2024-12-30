package types

import (
	"time"

	"gorm.io/gorm"
)

type Tour struct {
	Id             string `gorm:"primaryKey"`
	Source         string
	Destination    string
	StartDate      string
	EndDate        string
	Type           string
	Price          int32
	VehicleUnicode string
	TechnicalTeamID string
	TechnicalTeam  []*TechnicalTeam `gorm:"foreignKey:TechnicalTeamID"`
	CreatedAt      time.Time        `gorm:"autoCreateTime"` // Auto-set when created
	UpdatedAt      time.Time        `gorm:"autoUpdateTime"` // Auto-set when updated
	DeletedAt      gorm.DeletedAt   `gorm:"index"`          // Soft delete (optional)
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
