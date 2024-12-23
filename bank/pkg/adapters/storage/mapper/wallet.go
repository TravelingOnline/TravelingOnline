package mapper

import (
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

func WalletDomainToEntity(domainWallet *domain.Wallet) *types.Wallet {
	return &types.Wallet{
		UserID:  &domainWallet.UserID,
		Balance: domainWallet.Balance,
	}
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
	var toWl *types.Wallet
	fromWl := WalletDomainToEntity(domainTr.FromWallet)
	if !domainTr.IsPaidToSystem {
		toWl = WalletDomainToEntity(domainTr.ToWallet)
	}
	return &types.BankTransaction{
		Amount:         domainTr.Amount,
		FromWallet:     fromWl,
		ToWallet:       toWl,
		IsPaidToSystem: domainTr.IsPaidToSystem,
	}
}

func TransactionEntityToDomain(entity *types.BankTransaction) *domain.BankTransaction {
	var toWalDomain *domain.Wallet
	fromWalDomain := WalletEntityToDomain(entity.FromWallet)
	if !entity.IsPaidToSystem {
		toWalDomain = WalletEntityToDomain(entity.ToWallet)
	}
	return &domain.BankTransaction{
		Amount:         entity.Amount,
		Status:         entity.Status,
		FromWallet:     fromWalDomain,
		ToWallet:       toWalDomain,
		IsPaidToSystem: entity.IsPaidToSystem,
	}
}
func TransactionEntitiesToDomainTransactions(entities []types.BankTransaction) []domain.BankTransaction {
	var domainBankTransactions []domain.BankTransaction
	for _, e := range entities {
		domainBankTransactions = append(domainBankTransactions, domain.BankTransaction{Amount: e.Amount,
			Status:         e.Status,
			FromWallet:     WalletEntityToDomain(e.FromWallet),
			ToWallet:       WalletEntityToDomain(e.ToWallet),
			IsPaidToSystem: e.IsPaidToSystem,
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
