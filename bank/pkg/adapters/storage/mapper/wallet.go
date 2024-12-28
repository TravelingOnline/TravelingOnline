package mapper

import (
	"fmt"

	"github.com/onlineTraveling/bank/internal/bank/domain"
	"github.com/onlineTraveling/bank/pkg/adapters/storage/types"
)

func WalletEntityToDomain(entity *types.Wallet) *domain.Wallet {
	return &domain.Wallet{
		ID:      &entity.ID,
		UserID:  *entity.UserID,
		Balance: entity.Balance,
	}
}

func WalletDomainToEntity(domainWallet *domain.Wallet) (*types.Wallet, error) {
	if domainWallet == nil {

		return nil, fmt.Errorf("wallet is nil")
	}
	return &types.Wallet{
		// ID:      *domainWallet.ID, /////////////
		UserID:  &domainWallet.UserID,
		Balance: domainWallet.Balance,
	}, nil
}

func CreditCardEntityToDomain(entity *types.CreditCard) *domain.CreditCard {
	return &domain.CreditCard{
		ID:     entity.ID,
		Number: entity.Number,
	}
}

func CreditCardDomainToEntity(domainWallet *domain.CreditCard) *types.CreditCard {
	return &types.CreditCard{
		Number: domainWallet.Number,
	}
}

func DomainTransactionToTransactionEntity(domainTr *domain.BankTransaction) *types.BankTransaction {
	fw := &types.Wallet{
		ID:      *domainTr.FromWallet.ID,
		UserID:  &domainTr.FromWallet.UserID,
		Balance: domainTr.FromWallet.Balance,
	}

	tw := &types.Wallet{
		ID:      *domainTr.ToWallet.ID,
		UserID:  &domainTr.ToWallet.UserID,
		Balance: domainTr.ToWallet.Balance,
	}
	return &types.BankTransaction{
		Amount:       domainTr.Amount,
		FromWallet:   fw,
		ToWallet:     tw,
		FromWalletID: domainTr.FromWallet.ID,
		ToWalletID:   domainTr.ToWallet.ID,
		Status:       domainTr.Status,
	}
}

func TransactionEntityToDomain(entity *types.BankTransaction) *domain.BankTransaction {
	var toWalDomain *domain.Wallet
	fromWalDomain := WalletEntityToDomain(entity.FromWallet)
	toWalDomain = WalletEntityToDomain(entity.ToWallet)
	return &domain.BankTransaction{
		Amount:     entity.Amount,
		Status:     entity.Status,
		FromWallet: fromWalDomain,
		ToWallet:   toWalDomain,
	}
}
func TransactionEntitiesToDomainTransactions(entities []types.BankTransaction) []domain.BankTransaction {
	var domainBankTransactions []domain.BankTransaction
	for _, e := range entities {
		domainBankTransactions = append(domainBankTransactions, domain.BankTransaction{Amount: e.Amount,
			Status:     e.Status,
			FromWallet: WalletEntityToDomain(e.FromWallet),
			ToWallet:   WalletEntityToDomain(e.ToWallet),
		})
	}
	return domainBankTransactions
}

func BatchCreditCardEntityToDomain(entities []*types.CreditCard) []domain.CreditCard {
	var domainCreditCards []domain.CreditCard
	for _, e := range entities {
		domainCreditCards = append(domainCreditCards, domain.CreditCard{ID: e.ID, Number: e.Number})
	}
	return domainCreditCards
}