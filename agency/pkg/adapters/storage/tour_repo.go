package storage

import (
	"agency/internal/tour/domain"
	"agency/internal/tour/port"
	"agency/pkg/adapters/storage/mapper"
	"agency/pkg/adapters/storage/types"
	"context"

	"gorm.io/gorm"
)

type tourRepo struct {
	db *gorm.DB
}

func NewTourRepo(db *gorm.DB) port.Repo {
	return &tourRepo{db}
}

func (t *tourRepo) Create(ctx context.Context, tourDomain domain.Tour) error {
	tour := mapper.TourDomainToStorage(tourDomain)
	return t.db.Table("tours").WithContext(ctx).Create(tour).Error
}
func (t *tourRepo) GetAll(ctx context.Context, agencyID uint, page int, pageSize int) ([]domain.Tour, error) {
	tours := make([]domain.Tour, 0)

	q := t.db.WithContext(ctx).Table("tours").Debug()
	offset := (page - 1) * pageSize
	limit := pageSize

	err := q.Offset(offset).Limit(limit).Find(&tours).Error

	if err != nil {
		return nil, err
	}

	return tours, nil

}
func (t *tourRepo) GetByID(ctx context.Context, id uint) (*domain.Tour, error) {
	var tour types.Tour

	q := t.db.Table("tours").Debug().WithContext(ctx)

	q = q.Where("id = ?", id)

	err := q.First(&tour).Error

	if err != nil {
		return nil, err
	}
	return mapper.TourStorageToDomain(tour), nil
}
func (t *tourRepo) GetByFilter(ctx context.Context, filter *domain.TourFilter) (*domain.Tour, error) {
	var tour types.Tour

	q := t.db.Table("tours").Debug().WithContext(ctx)

	if filter.ID > 0 {
		q = q.Where("id = ? OR hotel_id = ? OR outbound_ticket_id = ? OR return_ticket_id = ?", filter.ID, filter.HotelID, filter.OutboundTicketID, filter.ReturnTicketID)
	}

	q = q.Where("is_active = ?", filter.IsActive)

	err := q.First(&tour).Error

	if err != nil {
		return nil, err
	}

	return mapper.TourStorageToDomain(tour), nil

}
func (t *tourRepo) Update(ctx context.Context, tour *domain.Tour) error {
	return nil
}
func (t *tourRepo) Delete(ctx context.Context, id uint) error {
	err := t.db.WithContext(ctx).Delete(&types.Tour{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
