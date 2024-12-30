package port

import (
	"context"

	"github.com/onlineTraveling/transport/internal/tour/domain"
)

type Repo interface {
	CreateTour(ctx context.Context, tour domain.Tour) (domain.TourID, error)
	UpdateTour(ctx context.Context, tour domain.Tour) (domain.TourID, error)
	DeleteTour(ctx context.Context, tourID domain.TourID) (domain.TourID, error)
	GetByIDTour(ctx context.Context, tourID domain.TourID) (domain.Tour, error)
}
