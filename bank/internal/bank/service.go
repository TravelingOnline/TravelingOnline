package bank

import (
	"context"

	"github.com/google/uuid"
	"github.com/onlineTraveling/bank/internal/bank/domain"
	"github.com/onlineTraveling/bank/internal/bank/port"
)

type WalletService struct {
	repo port.WalletRepo
}

func NewWalletService(repo port.WalletRepo) *WalletService {
	return &WalletService{
		repo: repo,
	}
}

func (o *WalletService) Create(ctx context.Context, wallet *domain.Wallet) (*domain.Wallet, error) {
	return o.repo.Create(ctx, wallet)
}

func (o *WalletService) Deposit(ctx context.Context, creditCard *domain.CreditCard, amount uint, userID uuid.UUID) (*domain.Wallet, error) {
	return o.repo.Deposit(ctx, creditCard, amount, userID)
}

func (o *WalletService) Withdraw(ctx context.Context, creditCard *domain.CreditCard, amount uint, userID uuid.UUID) (*domain.Wallet, error) {
	return o.repo.Withdraw(ctx, creditCard, amount, userID)
}

func (o *WalletService) GetWallet(ctx context.Context, userID uuid.UUID) (*domain.Wallet, error) {
	return o.repo.GetWallet(ctx, userID)
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

//////////////////////////////

type BankTransactionService struct {
	repo port.BankTransactionRepo
}

func NewBankTransactionService(repo port.BankTransactionRepo) *BankTransactionService {
	return &BankTransactionService{
		repo: repo,
	}
}
func (b *BankTransactionService) Transfer(ctx context.Context, transaction *domain.BankTransaction) (*domain.BankTransaction, error) {
	return b.repo.Transfer(ctx, transaction)
}
