package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tour struct {
	gorm.Model
	UUID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	OutboundTicketID uint
	ReturnTicketID   uint
	AgencyID         uint
	HotelID          uint
	Capacity         uint
	Price            uint64
	IsActive         bool
}
