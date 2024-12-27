package storage

import (
	"context"
	// "errors"

	"github.com/google/uuid"
	"github.com/onlineTraveling/bank/internal/bank/domain"
	"github.com/onlineTraveling/bank/internal/bank/port"
	"github.com/onlineTraveling/bank/pkg/adapters/storage/mapper"
	"github.com/onlineTraveling/bank/pkg/adapters/storage/types"

	// _context "github.com/onlineTraveling/bank/pkg/context"

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
	// Begin a transaction
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	var walletFrom, walletTo *domain.Wallet

	// Fetch the source wallet
	if err := tx.Table("wallets").Where("id=?", tr.FromWallet.ID).First(&walletFrom).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Fetch the destination wallet
	if err := tx.Table("wallets").Where("id=?", tr.ToWallet.ID).First(&walletTo).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Check transfer amount and balance
	if tr.Amount < 100 || walletFrom.Balance < tr.Amount {
		tx.Rollback()
		return nil, port.ErrNotEnoughBalance
	}

	// Deduct from source wallet
	walletFrom.Balance -= tr.Amount
	if err := tx.WithContext(ctx).Save(&walletFrom).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Fetch the system wallet
	var systemWalEntity *types.Wallet
	if err := tx.Where("is_system_wallet = ?", true).First(&systemWalEntity).Error; err != nil {
		tx.Rollback()
		return nil, port.ErrSystemWalletDoesNotExists
	}

	// Fetch commission settings
	var commissionEntity *types.Commission
	if err := tx.WithContext(ctx).First(&commissionEntity).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Calculate and apply commission
	commission := tr.Amount * commissionEntity.AppCommissionPercentage / 100
	walletTo.Balance += (tr.Amount - commission)

	// Save destination wallet
	if err := tx.WithContext(ctx).Save(&walletTo).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Update system wallet balance
	systemWalEntity.Balance += commission
	if err := tx.WithContext(ctx).Save(&systemWalEntity).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Update transaction status
	tr.Status = types.TransactionSuccess
	createdTransaction := mapper.DomainTransactionToTransactionEntity(tr)
	if err := tx.WithContext(ctx).Create(&createdTransaction).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return tr, nil
}
