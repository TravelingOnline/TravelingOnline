package port

import (
	"context"

	"github.com/onlineTraveling/auth/internal/codeVerification/domain"
	"github.com/onlineTraveling/auth/internal/common"
	userDomain "github.com/onlineTraveling/auth/internal/user/domain"
)

type Repo interface {
	Create(ctx context.Context, codeVerification *domain.CodeVerification) (domain.CodeVerificationID, error)
	CreateOutbox(ctx context.Context, outbox *domain.CodeVerificationOutbox) error
	QueryOutboxes(ctx context.Context, limit uint, status common.OutboxStatus) ([]domain.CodeVerificationOutbox, error)
	GetUserCodeVerificationValue(ctx context.Context, userID userDomain.UserID) (string, error)
}
