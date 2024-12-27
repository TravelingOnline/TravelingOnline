package storage

import (
	"agency/internal/tour/domain"
	"agency/internal/tour/port"
	"context"

	"gorm.io/gorm"
)

type tourRepo struct {
	db *gorm.DB
}

func NewTourRepo(db *gorm.DB) port.Repo {
	return &tourRepo{db}
}

func (t *tourRepo) Create(ctx context.Context, tour domain.Tour) error {
	return nil
}
func (t *tourRepo) GetAll(ctx context.Context, agencyID uint, page int, pagesize int) ([]domain.Tour, error) {
	return nil, nil
}
func (t *tourRepo) GetByID(ctx context.Context, id uint) (*domain.Tour, error) {
	return nil, nil
}
func (t *tourRepo) GetByFilter(ctx context.Context, filter *domain.TourFilter) (*domain.Tour, error) {
	return nil, nil
}
func (t *tourRepo) Update(ctx context.Context, tour *domain.Tour) error {
	return nil
}
func (t *tourRepo) Delete(ctx context.Context, id uint) error {
	return nil
}
