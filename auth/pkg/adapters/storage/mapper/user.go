package mapper

import (
	"github.com/onlineTraveling/auth/internal/user/domain"
	"github.com/onlineTraveling/auth/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

func UserDomain2Storage(userDomain domain.User) *types.User {

	return &types.User{
		Model: gorm.Model{
			ID:        uint(userDomain.ID),
			CreatedAt: userDomain.CreatedAt,
			DeletedAt: gorm.DeletedAt(ToNullTime(userDomain.DeletedAt)),
			UpdatedAt: userDomain.UpdatedAt,
		},
		Email:        string(userDomain.Email),
		PasswordHash: userDomain.PasswordHash,
	}
}

func UserStorage2Domain(user types.User) *domain.User {

	return &domain.User{
		ID:        domain.UserID(user.ID),
		CreatedAt: user.CreatedAt,
		// DeletedAt:         user.DeletedAt,
		UpdatedAt:    user.UpdatedAt,
		Email:        domain.Email(user.Email),
		PasswordHash: user.PasswordHash,
	}
}
