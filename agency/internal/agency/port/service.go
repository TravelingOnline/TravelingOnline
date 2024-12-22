package port

import (
	"agency/internal/agency/domain"
	"context"
)

type Service interface {
	CreateAgency(ctx context.Context, agency domain.Agency) error
	GetAllAgencies(ctx context.Context, page int, pageSize int) ([]domain.Agency, error)
	GetAgencyByID(ctx context.Context, id uint) (*domain.Agency, error)
	GetAgencyByOwnerID(ctx context.Context, id uint) (*domain.Agency, error)
	UpdateAgency(ctx context.Context, agency *domain.Agency) error
	DeleteAgency(ctx context.Context, id uint) error
}
