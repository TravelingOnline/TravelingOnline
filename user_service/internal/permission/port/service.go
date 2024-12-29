package port

import (
	"context"
	"user_service/internal/permission/domain"
)

//TODO

type Service interface {
	CreatePermision(ctx context.Context, permision domain.Permision) (domain.PermisionID, error)
	GetPermisionById(ctx context.Context, permisionId domain.PermisionID) (*domain.Permision, error)
	GetPermisionByName(ctx context.Context, name string) (*domain.Permision, error)
}
