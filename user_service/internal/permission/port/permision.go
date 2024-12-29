package port

import (
	"context"
	"user_service/internal/permission/domain"
)

type Repo interface {
	Create(ctx context.Context, permision domain.Permision) (domain.PermisionID, error)
	GetById(ctx context.Context, permisionId domain.PermisionID) (*domain.Permision, error)
	GetByName(ctx context.Context, name string) (*domain.Permision, error)
}
