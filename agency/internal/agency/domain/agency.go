package domain

import (
	"time"
)

type (
	AgencyID uint
	OwnerID  uint
)

type Agency struct {
	ID        AgencyID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	OwnerID   OwnerID
}
