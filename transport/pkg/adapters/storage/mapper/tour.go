package mapper

import (
	"errors"
	"time"

	"github.com/onlineTraveling/transport/internal/tour/domain"
	"github.com/onlineTraveling/transport/pkg/adapters/storage/types"
	"gorm.io/gorm"
)

func TourStroage2Domain(t types.Tour) (domain.Tour, error) {
	// Check if the required fields are valid
	if t.Id == "" {
		return domain.Tour{}, errors.New("tour ID is required")
	}
	if t.TechnicalTeam == nil {
		return domain.Tour{}, errors.New("Techteam information is missing")
	}

	// Construct the domain company
	company := domain.Tour{
		Id:             t.Id,
		Source:         t.Source,
		Destination:    t.Destination,
		StartDate:      t.StartDate,
		EndDate:        t.EndDate,
		Type:           t.Type,
		Price:          t.Price,
		VehicleUnicode: t.VehicleUnicode,
		TechnicalTeam:  []*domain.TechnicalTeam{},
	}

	// Return the constructed company and nil error if no validation failed
	return company, nil
}

func DomainTour2Storage(t domain.Tour) types.Tour {
	return types.Tour{
		Id:             t.Id,
		Source:         t.Source,
		Destination:    t.Destination,
		StartDate:      t.StartDate,
		EndDate:        t.EndDate,
		Type:           t.Type,
		Price:          t.Price,
		VehicleUnicode: t.VehicleUnicode,
		TechnicalTeam:  []*types.TechnicalTeam{},
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      gorm.DeletedAt{},
	}
}
