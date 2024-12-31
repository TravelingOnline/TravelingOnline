package port

import (
	"context"

	"github.com/onlineTraveling/transport/internal/tour/domain"
)

type Service interface {
	CreateTourService(ctx context.Context, tour domain.Tour) (domain.TourID, error)
	UpdateTourService(ctx context.Context, tour domain.Tour) (domain.TourID, error)
	DeleteTourService(ctx context.Context, tourID domain.TourID) (domain.TourID, error)
	GetByIDTourService(ctx context.Context, tourID domain.TourID) (domain.Tour, error)

}
