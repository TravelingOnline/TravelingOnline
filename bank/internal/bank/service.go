package bank

import (
	"context"

	"github.com/google/uuid"
	"github.com/onlineTraveling/bank/internal/bank/domain"
	"github.com/onlineTraveling/bank/internal/bank/port"
)

type walletService struct {
	repo port.WalletRepo
}

func NewWalletService(repo port.WalletRepo) *walletService {
	return &walletService{
		repo: repo,
	}
}

func (o *walletService) Create(ctx context.Context, wallet *domain.Wallet) (*domain.Wallet, error) {
	return o.repo.Create(ctx, wallet)
}

func (o *walletService) Deposit(ctx context.Context, creditCard *domain.CreditCard, amount uint, userID uuid.UUID) (*domain.Wallet, error) {
	return o.repo.Deposit(ctx, creditCard, amount, userID)
}

func (o *walletService) Withdraw(ctx context.Context, creditCard *domain.CreditCard, amount uint, userID uuid.UUID) (*domain.Wallet, error) {
	return o.repo.Withdraw(ctx, creditCard, amount, userID)
}

func (o *walletService) GetWallet(ctx context.Context, userID uuid.UUID) (*domain.Wallet, error) {
	return o.repo.GetWallet(ctx, userID)
}

func (o *walletService) Transfer(ctx context.Context, transaction *domain.BankTransaction) (*domain.BankTransaction, error) {
	return o.repo.Transfer(ctx, transaction)
}

////////////////////////////////////////////////////////

type CreditCardService struct {
	repo port.CreditCardRepo
}

func NewCreditCardService(repo port.CreditCardRepo) *CreditCardService {
	return &CreditCardService{
		repo: repo,
	}
}

func (o *CreditCardService) CreateCardAndAddToWallet(ctx context.Context, creditCard *domain.CreditCard, userID uuid.UUID) (*domain.CreditCard, error) {
	if !domain.IsValidCardNumber(creditCard.Number) {
		return nil, port.ErrInvalidCardNumber
	}
	return o.repo.CreateCardAndAddToWallet(ctx, creditCard, userID)
}

func (o *CreditCardService) GetUserWalletCards(ctx context.Context, userID uuid.UUID) ([]domain.CreditCard, error) {
	return o.repo.GetUserWalletCards(ctx, userID)
}
