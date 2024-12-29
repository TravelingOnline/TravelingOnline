package mapper

import (
	"time"
	roleDomain "user_service/internal/role/domain"
	"user_service/internal/user/domain"
	"user_service/pkg/adapters/storage/entities"

	"gorm.io/gorm"
)

func MapInternalUserToEntity(internalUser domain.User) entities.User {
	entityRoles := make([]*entities.Role, 0)
	if internalUser.Role.ID != 0 {
		entityRoles = append(entityRoles, &entities.Role{
			ID:   uint(internalUser.Role.ID),
			Name: internalUser.Role.Name,
		})
	}

	return entities.User{
		ID:           uint(internalUser.ID),
		FullName:     internalUser.FullName,
		Email:        internalUser.Email,
		Password:     internalUser.Password,
		NationalCode: internalUser.NationalCode,
		Roles:        entityRoles,
		CreatedAt:    internalUser.CreatedAt,
		UpdatedAt:    time.Now(), // UpdatedAt is not in the internal User, so use the current time.
		DeletedAt:    gorm.DeletedAt{Time: internalUser.DeletedAt, Valid: !internalUser.DeletedAt.IsZero()},
	}
}

func MapEntityUserToInternal(entityUser entities.User) domain.User {
	var internalRole roleDomain.Role
	if len(entityUser.Roles) > 0 {
		role := entityUser.Roles[0]
		internalRole = roleDomain.Role{
			ID:   roleDomain.RoleId(role.ID),
			Name: role.Name,
		}
	}

	return domain.User{
		ID:           domain.UserId(entityUser.ID),
		FullName:     entityUser.FullName,
		Email:        entityUser.Email,
		Password:     entityUser.Password,
		NationalCode: entityUser.NationalCode,
		CreatedAt:    entityUser.CreatedAt,
		DeletedAt:    entityUser.DeletedAt.Time,
		Role:         internalRole,
	}
}
