package storage

import (
	"agency/internal/agency/domain"
	"agency/internal/agency/port"
	"agency/pkg/adapters/storage/mapper"
	"agency/pkg/adapters/storage/types"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type agencyRepo struct {
	db *gorm.DB
}

func NewAgencyRepo(db *gorm.DB) port.Repo {
	return &agencyRepo{db}
}

func (r *agencyRepo) Create(ctx context.Context, agencyDomain domain.Agency) error {
	agency := mapper.AgencyDomainToStorage(agencyDomain)
	return r.db.Table("agencies").WithContext(ctx).Create(agency).Error
}

func (r *agencyRepo) GetAll(ctx context.Context, page int, pageSize int) ([]domain.Agency, error) {

	agencies := make([]domain.Agency, 0)

	q := r.db.WithContext(ctx).Table("agencies").Debug()

	offset := (page - 1) * pageSize
	limit := pageSize
	err := q.Offset(offset).Limit(limit).Find(&agencies).Error

	if err != nil {
		return nil, err
	}
	return agencies, nil
}

func (r *agencyRepo) GetByID(ctx context.Context, id uint) (*domain.Agency, error) {
	var agency types.Agency

	q := r.db.Table("agencies").Debug().WithContext(ctx)

	q = q.Where("id = ?", id)

	err := q.First(&agency).Error

	if err != nil {
		return nil, err
	}

	return mapper.AgencyStorageToDomain(agency), nil
}

func (r *agencyRepo) GetByOwnerID(ctx context.Context, id uint) (*domain.Agency, error) {
	var agency types.Agency

	q := r.db.Table("agencies").Debug().WithContext(ctx)

	q = q.Where("owner_id = ?", id)

	err := q.First(&agency).Error

	if err != nil {
		return nil, err
	}

	return mapper.AgencyStorageToDomain(agency), nil
}

func (r *agencyRepo) Update(ctx context.Context, agencyDomain *domain.Agency) error {
	if agencyDomain == nil {
		return errors.New("agency cannot be nil")
	}

	agency := mapper.AgencyDomainToStorage(*agencyDomain)

	err := r.db.WithContext(ctx).Where("id = ?", agency.ID).Updates(agency).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *agencyRepo) Delete(ctx context.Context, id uint) error {

	err := r.db.WithContext(ctx).Delete(&types.Agency{}, id).Error
	fmt.Println(err)
	if err != nil {
		return err
	}

	return nil
}
