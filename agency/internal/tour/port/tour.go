package port

import (
	"agency/internal/tour/domain"
	"context"
)

type Repo interface {
	Create(ctx context.Context, tour domain.Tour) error
	GetAll(ctx context.Context, agencyID uint, page int, pagesize int) ([]domain.Tour, error)
	GetByID(ctx context.Context, id uint) (*domain.Tour, error)
	GetByFilter(ctx context.Context, filter *domain.TourFilter) (*domain.Tour, error)
	Update(ctx context.Context, tour *domain.Tour) error
	Delete(ctx context.Context, id uint) error
}
