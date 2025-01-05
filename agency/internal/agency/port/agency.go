package port

import (
	"agency/internal/agency/domain"
	"context"
)

type Repo interface {
	Create(ctx context.Context, agency domain.Agency) error
	GetAll(ctx context.Context, page int, pageSize int) ([]domain.Agency, error)
	GetByID(ctx context.Context, id uint) (*domain.Agency, error)
	GetByOwnerID(ctx context.Context, id uint) (*domain.Agency, error)
	Update(ctx context.Context, agency *domain.Agency) error
	Delete(ctx context.Context, id uint) error
}
