package domain

import (
	"time"

	"github.com/google/uuid"
)

type (
	TourID   uint
	TourUUID uuid.UUID
)

type Tour struct {
	ID               TourID
	UUID             TourUUID
	CreatedAt        time.Time
	UpdatedAt        time.Time
	OutboundTicketID uint
	ReturnTicketID   uint
	AgencyID         uint
	HotelID          uint
	Capacity         uint
	Price            uint64
	IsActive         bool
}

type TourFilter struct {
	ID               TourID
	HotelID          uint
	OutboundTicketID uint
	ReturnTicketID   uint
	IsActive         bool
}
