package mapper

import (
	"user_service/internal/role/domain"
	"user_service/pkg/adapters/storage/entities"
)

// TODO
func MapInternalRoleToEntity(domainRole *domain.Role) entities.Role {
	return entities.Role{
		ID:   uint(domainRole.ID),
		Name: domainRole.Name,
	}
}
