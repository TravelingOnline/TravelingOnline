package port

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/onlineTraveling/bank/internal/bank/domain"
)

var (
	ErrNotEnoughBalance          = errors.New("not enough balance")
	ErrUserAlreadyHasWallet      = errors.New("user already has wallet")
	ErrUserWalletDoesNotExists   = errors.New("user wallet does not exists")
	ErrSystemWalletDoesNotExists = errors.New("system wallet does not exists")
	ErrMinTrans                  = errors.New("minimum value of transaction is 100")
)

type WalletRepo interface {
	Create(ctx context.Context, user *domain.Wallet) (*domain.Wallet, error)
	Deposit(ctx context.Context, creditCard *domain.CreditCard, amount uint64, userID uuid.UUID) (*domain.Wallet, error)
	Withdraw(ctx context.Context, creditCard *domain.CreditCard, amount uint64, userID uuid.UUID) (*domain.Wallet, error)
	GetWallet(ctx context.Context, userID uuid.UUID) (*domain.Wallet, error)
	DeleteWallet(ctx context.Context, userID uuid.UUID) error
}
