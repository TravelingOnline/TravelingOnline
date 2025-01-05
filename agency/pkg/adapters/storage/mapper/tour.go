package mapper

import (
	"agency/internal/tour/domain"
	"agency/pkg/adapters/storage/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func TourDomainToStorage(tourDomain domain.Tour) *types.Tour {
	return &types.Tour{
		Model: gorm.Model{
			ID:        uint(tourDomain.ID),
			CreatedAt: tourDomain.CreatedAt,
			UpdatedAt: tourDomain.UpdatedAt,
		},
		UUID:             uuid.UUID(tourDomain.UUID),
		OutboundTicketID: tourDomain.OutboundTicketID,
		ReturnTicketID:   tourDomain.ReturnTicketID,
		AgencyID:         tourDomain.AgencyID,
		HotelID:          tourDomain.HotelID,
		Capacity:         tourDomain.Capacity,
		Price:            tourDomain.Price,
		IsActive:         tourDomain.IsActive,
	}
}

func TourStorageToDomain(tour types.Tour) *domain.Tour {
	return &domain.Tour{
		ID:               domain.TourID(tour.ID),
		CreatedAt:        tour.CreatedAt,
		UpdatedAt:        tour.UpdatedAt,
		UUID:             domain.TourUUID(tour.UUID),
		OutboundTicketID: tour.OutboundTicketID,
		ReturnTicketID:   tour.ReturnTicketID,
		AgencyID:         tour.AgencyID,
		HotelID:          tour.HotelID,
		Capacity:         tour.Capacity,
		Price:            tour.Price,
		IsActive:         tour.IsActive,
	}
}
