package storage

import (
	"context"

	"github.com/google/uuid"
	"github.com/onlineTraveling/bank/internal/bank/domain"
	"github.com/onlineTraveling/bank/internal/bank/port"
	"github.com/onlineTraveling/bank/pkg/adapters/storage/mapper"
	"github.com/onlineTraveling/bank/pkg/adapters/storage/types"
	"gorm.io/gorm"
)

type bankTransactionRepo struct {
	db *gorm.DB
}

func NewBankTransactionRepo(db *gorm.DB) port.BankTransactionRepo {
	return &bankTransactionRepo{
		db: db,
	}
}
func (r *bankTransactionRepo) GetTransactionsByUserId(ctx context.Context, userID uuid.UUID) ([]domain.BankTransaction, error) {
	var transactions []types.BankTransaction
	err := r.db.Where("user_id =?", userID).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return mapper.TransactionEntitiesToDomainTransactions(transactions), nil
}
func (r *bankTransactionRepo) Transfer(ctx context.Context, tr *domain.BankTransaction) (*domain.BankTransaction, error) {
	transaction := mapper.DomainTransactionToTransactionEntity(tr)
	var wallets []types.Wallet
	var walletIDs []uuid.UUID
	if tr.IsPaidToSystem {
		walletIDs = []uuid.UUID{*transaction.FromWallet.UserID}
	} else {
		walletIDs = []uuid.UUID{*transaction.FromWallet.UserID, *transaction.ToWallet.UserID}
	}
	err := r.db.Where("user_id IN ?", walletIDs).Find(&wallets).Error
	if err != nil {
		return nil, err
	}
	if transaction.Amount < 100 {
		return nil, port.ErrNotEnoughBalance
	}
	if !(len(wallets) == 2 && !transaction.IsPaidToSystem || len(wallets) == 1 && transaction.IsPaidToSystem) {
		return nil, port.ErrUserWalletDoesNotExists
	}
	var fromWalEntity *types.Wallet
	var toWalEntity *types.Wallet
	for _, wal := range wallets {
		if *wal.UserID == *transaction.FromWallet.UserID {
			fromWalEntity = &wal
			transaction.FromWallet.ID = fromWalEntity.ID
		} else if *wal.UserID == *transaction.ToWallet.UserID {
			toWalEntity = &wal
			transaction.ToWallet.ID = toWalEntity.ID
		}
	}
	if fromWalEntity.Balance < transaction.Amount {
		return nil, port.ErrNotEnoughBalance
	}
	fromWalEntity.Balance -= transaction.Amount
	if err := r.db.WithContext(ctx).Save(&fromWalEntity).Error; err != nil {
		return nil, err
	}
	var systemWalEntity *types.Wallet
	err = r.db.Where("is_system_wallet = ?", true).First(&systemWalEntity).Error
	if err != nil {
		return nil, port.ErrSystemWalletDoesNotExists
	}
	if transaction.IsPaidToSystem {
		systemWalEntity.Balance += transaction.Amount
	} else {
		var commissionEntity *types.Commission
		err := r.db.WithContext(ctx).First(&commissionEntity).Error
		if err != nil {
			return nil, err
		}
		tax := transaction.Amount * commissionEntity.AppCommissionPercentage / 100
		toWalEntity.Balance += transaction.Amount - tax
		if err := r.db.WithContext(ctx).Save(&toWalEntity).Error; err != nil {
			return nil, err
		}
		systemWalEntity.Balance += tax
	}
	if err := r.db.WithContext(ctx).Save(&systemWalEntity).Error; err != nil {
		return nil, err
	}
	transaction.Status = types.TransactionSuccess
	err = r.db.WithContext(ctx).Create(&transaction).Error
	if err != nil {
		return nil, err
	}
	createdTransaction := mapper.TransactionEntityToDomain(transaction)
	return createdTransaction, nil
}
