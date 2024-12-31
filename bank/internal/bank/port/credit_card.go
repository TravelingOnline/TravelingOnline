package port

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/onlineTraveling/bank/internal/bank/domain"
)

var (
	ErrInvalidCardNumber = errors.New("invalid card number")
	ErrCardAlreadyExists = errors.New("card already exists in wallet")
)

type CreditCardRepo interface {
	CreateCardAndAddToWallet(ctx context.Context, creditCard *domain.CreditCard, userID uuid.UUID) (*domain.CreditCard, error)
	GetUserWalletCards(ctx context.Context, userID uuid.UUID) ([]domain.CreditCard, error)
	UpdateCard(ctx context.Context, creditCard *domain.CreditCard) error
	DeleteCard(ctx context.Context, creditCardId uuid.UUID) error
}

// func NewCreditCard(number string) *domain.CreditCard {
// 	return &domain.CreditCard{
// 		Number: number,
// 	}
// }
