package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/onlineTraveling/bank/internal/bank"
	"github.com/onlineTraveling/bank/internal/bank/domain"
)

type BankService struct {
	walletService          *bank.WalletService
	creditCardService      *bank.CreditCardService
	bankTransactionService *bank.BankTransactionService
}

func NewBankService(walletService *bank.WalletService, creditCardService *bank.CreditCardService, bankTransactionService *bank.BankTransactionService) *BankService {
	return &BankService{
		walletService:          walletService,
		creditCardService:      creditCardService,
		bankTransactionService: bankTransactionService,
	}
}

func (s *BankService) CreateWallet(ctx context.Context, wl *domain.Wallet) (*domain.Wallet, error) {
	createdWallet, err := s.walletService.Create(ctx, wl)
	if err != nil {
		return nil, err
	}
	return createdWallet, nil
}

func (s *BankService) AddCardToWalletByUserID(ctx context.Context, card *domain.CreditCard, userID uuid.UUID) (*domain.CreditCard, error) {
	createdCard, err := s.creditCardService.CreateCardAndAddToWallet(ctx, card, userID)
	if err != nil {
		return nil, err
	}
	return createdCard, nil
}

func (s *BankService) GetUserWalletCards(ctx context.Context, userID uuid.UUID) ([]domain.CreditCard, error) {
	userWalletCards, err := s.creditCardService.GetUserWalletCards(ctx, userID)
	if err != nil {
		return nil, err
	}
	return userWalletCards, nil
}

func (s *BankService) Deposit(ctx context.Context, card *domain.CreditCard, amount uint, userID uuid.UUID) (*domain.Wallet, error) {
	userWallet, err := s.walletService.Deposit(ctx, card, amount, userID)
	if err != nil {
		return nil, err
	}
	return userWallet, nil
}

func (s *BankService) Withdraw(ctx context.Context, card *domain.CreditCard, amount uint, userID uuid.UUID) (*domain.Wallet, error) {
	userWallet, err := s.walletService.Withdraw(ctx, card, amount, userID)
	if err != nil {
		return nil, err
	}
	return userWallet, nil
}

func (s *BankService) GetWallet(ctx context.Context, userID uuid.UUID) (*domain.Wallet, error) {
	userWallet, err := s.walletService.GetWallet(ctx, userID)
	if err != nil {
		return nil, err
	}
	return userWallet, nil
}

func (s *BankService) Transfer(ctx context.Context, tr *domain.BankTransaction) (*domain.BankTransaction, error) {
	createdTransaction, err := s.bankTransactionService.Transfer(ctx, tr)
	if err != nil {
		return nil, err
	}
	return createdTransaction, nil
}
