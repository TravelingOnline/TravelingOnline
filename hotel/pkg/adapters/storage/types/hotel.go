package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Hotel struct {
	gorm.Model
	OwnerID   uuid.UUID `gorm:"uniqueIndex:idx_owner_name"`
	Name      string    `gorm:"uniqueIndex:idx_owner_name"`
	City      string
	Country   string
	Address   string
	Rate      int32
	Details   string
	IsBlocked bool
	Rooms     []Room
}
