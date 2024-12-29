package port

import (
	"context"
	"user_service/internal/user/domain"
)

type Repo interface {
	Create(ctx context.Context, user domain.User) (domain.UserId, error)
	GetById(ctx context.Context, userId domain.UserId) (*domain.User, error)
}
