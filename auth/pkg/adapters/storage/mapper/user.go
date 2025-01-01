package mapper

import (
	"github.com/onlineTraveling/auth/internal/user/domain"
	"github.com/onlineTraveling/auth/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

func UserDomain2Storage(userDomain domain.User) *types.User {

	return &types.User{
		ID:        userDomain.ID,
		CreatedAt: userDomain.CreatedAt,
		DeletedAt: gorm.DeletedAt(ToNullTime(userDomain.DeletedAt)),
		UpdatedAt: userDomain.UpdatedAt,

		Email:        string(userDomain.Email),
		PasswordHash: userDomain.Password,
	}
}

func UserStorage2Domain(user types.User) *domain.User {

	return &domain.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		// DeletedAt:         user.DeletedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     domain.Email(user.Email),
		Password:  user.PasswordHash,
	}
}
