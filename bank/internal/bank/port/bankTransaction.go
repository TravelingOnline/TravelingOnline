package port

import (
	"context"
	"errors"

	// "github.com/google/uuid"
	"github.com/onlineTraveling/bank/internal/bank/domain"
)

var (
	ErrInvalidUserId = errors.New("invalid user id")
	ErrNotTransfered = errors.New("transfer not take action")
)

type BankTransactionRepo interface {
	// GetTransactionsByUserId(ctx context.Context, userID uuid.UUID) ([]domain.BankTransaction, error)
	Transfer(ctx context.Context, transaction *domain.BankTransaction) (*domain.BankTransaction, error)
}
