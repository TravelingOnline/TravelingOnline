package port

import (
	"context"
	"user_service/internal/role/domain"
)

type Service interface {
	CreateRole(ctx context.Context, role *domain.Role) (domain.RoleId, error)
	GetRoleById(ctx context.Context, roleId domain.RoleId) (*domain.Role, error)
	GetRoleByName(ctx context.Context, roleName string) (*domain.Role, error)
}
