package port

import (
	"context"

	"github.com/onlineTraveling/auth/internal/user/domain"
	"github.com/onlineTraveling/auth/pkg/jwt"
)

type Service interface {
	CreateUser(ctx context.Context, user domain.User) (domain.UserID, error)
	GetUserByID(ctx context.Context, userID domain.UserID) (*domain.User, error)
	GetUserIDFromToken(ctx context.Context, Token string) (*jwt.UserClaims, error)
	GetUserByEmail(ctx context.Context, email domain.Email) (*domain.User, error)

	UpdateUser(ctx context.Context, user domain.User) error
	DeleteByID(ctx context.Context, userID domain.UserID) error
	GetUserByFilter(ctx context.Context, filter *domain.UserFilter) (*domain.User, error)
}
