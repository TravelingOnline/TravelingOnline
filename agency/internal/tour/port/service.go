package port

import (
	"agency/internal/tour/domain"
	"context"
)

type Service interface {
	CreateTour(ctx context.Context, tour domain.Tour) error
	GetAllTours(ctx context.Context, agencyID uint, page int, pagesize int) ([]domain.Tour, error)
	GetTourByID(ctx context.Context, id uint) (*domain.Tour, error)
	UpdateTour(ctx context.Context, tour *domain.Tour) error
	DeleteTour(ctx context.Context, id uint) error
	GetTourByFilter(ctx context.Context, filter *domain.TourFilter) (*domain.Tour, error)
}
