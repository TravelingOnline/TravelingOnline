package port

import (
	"context"
	"user_service/internal/role/domain"
)

type Repo interface {
	Create(ctx context.Context, role domain.Role) (domain.RoleId, error)
	GetById(ctx context.Context, roleId domain.RoleId) (*domain.Role, error)
	GetByName(ctx context.Context, roleName string) (*domain.Role, error)
}
