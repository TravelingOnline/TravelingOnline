package domain

import (
	"time"

	"github.com/google/uuid"
)

type (
	AgencyID uint
	OwnerID  uuid.UUID
)

type Agency struct {
	ID        AgencyID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	OwnerID   OwnerID
}
