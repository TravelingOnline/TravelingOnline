package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/onlineTraveling/bank/internal/bank/domain"
)

type WalletCreditCardRepo interface {
	Create(ctx context.Context, creditCard *domain.WalletCreditCard) (*domain.WalletCreditCard, error)
}

func NewWalletCreditCard(walletID uuid.UUID, creditCardID uuid.UUID) *domain.WalletCreditCard {
	return &domain.WalletCreditCard{WalletID: walletID, CreditCardID: creditCardID}
}
