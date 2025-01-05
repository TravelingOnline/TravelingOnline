package port

import (
	"context"

	"github.com/onlineTraveling/auth/internal/codeVerification/domain"
	"github.com/onlineTraveling/auth/internal/common"
	userDomain "github.com/onlineTraveling/auth/internal/user/domain"
)

type Service interface {
	Send(ctx context.Context, codeVerification *domain.CodeVerification) error
	CheckUserCodeVerificationValue(ctx context.Context, userID userDomain.UserID, val string) (bool, error)
	common.OutboxHandler[domain.CodeVerificationOutbox]
}
