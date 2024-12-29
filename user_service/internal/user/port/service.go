package port

import (
	"context"
	"user_service/internal/user/domain"
)

type Service interface {
	CreateUser(ctx context.Context, user domain.User) (domain.UserId, error)
	GetUserById(ctx context.Context, userId domain.UserId) (*domain.User, error)
}
