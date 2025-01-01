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

	// Construct the domain company
	company := domain.Tour{
		Id:          t.Id,
		Source:      t.Source,
		Destination: t.Destination,
		StartDate:   t.StartDate,
		EndDate:     t.EndDate,
		Type:        t.Type,
		Price:       t.Price,
		CompanyID:   t.CompanyID,
		Vehicle: domain.Vehicle{
			Id:              t.Vehicle.Id,
			Unicode:         t.Vehicle.Unicode,
			RequiredExperts: t.Vehicle.RequiredExperts,
			Speed:           t.Vehicle.Speed,
			RentPrice:       t.Vehicle.RentPrice,
			Type:            t.Vehicle.Type,
			Passenger:       t.Vehicle.Passenger,
			Model:           t.Vehicle.Model,
		},
		// TechnicalTeam: []*domain.TechnicalTeam{},
	}

	// Return the constructed company and nil error if no validation failed
	return company, nil
}

func DomainTour2Storage(t domain.Tour) types.Tour {
	// Map TechnicalTeam from domain to storage
	technicalTeams := make([]*types.TechnicalTeam, len(t.TechnicalTeam))
	for i, team := range t.TechnicalTeam {
		technicalTeams[i] = &types.TechnicalTeam{
			Id:        team.Id,
			FirstName: team.FirstName,
			LastName:  team.LastName,
			Age:       team.Age,
			Expertise: team.Expertise,
		}
	}

	// Map Tour
	return types.Tour{
		Id:          t.Id,
		Source:      t.Source,
		Destination: t.Destination,
		StartDate:   t.StartDate,
		EndDate:     t.EndDate,
		Type:        t.Type,
		Price:       t.Price,
		CompanyID:   t.CompanyID,
		Vehicle: &types.Vehicle{
			Id:              t.Vehicle.Id,
			Unicode:         t.Vehicle.Unicode,
			RequiredExperts: t.Vehicle.RequiredExperts,
			Speed:           t.Vehicle.Speed,
			RentPrice:       t.Vehicle.RentPrice,
			Type:            t.Vehicle.Type,
			Passenger:       t.Vehicle.Passenger,
			Model:           t.Vehicle.Model,
			CreatedAt:       time.Time{},
			UpdatedAt:       time.Time{},
			DeletedAt:       gorm.DeletedAt{},
		},
		TechnicalTeam: technicalTeams,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     gorm.DeletedAt{},
	}
}
