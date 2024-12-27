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
	var walletFrom *domain.Wallet
	var walletTo *domain.Wallet

	er := r.db.Table("wallets").Where("id=?", tr.FromWallet.ID).First(&walletFrom).Error
	if er != nil {
		return nil, er
	}
	er = r.db.Table("wallets").Where("id=?", tr.ToWallet.ID).First(&walletTo).Error
	if er != nil {
		return nil, er
	}

	if tr.Amount < 100 {
		return nil, port.ErrNotEnoughBalance
	}

	if walletFrom.Balance < tr.Amount {
		return nil, port.ErrNotEnoughBalance
	}
	walletFrom.Balance -= tr.Amount
	if err := r.db.WithContext(ctx).Save(&walletFrom).Error; err != nil {
		return nil, err
	}
	var systemWalEntity *types.Wallet
	err := r.db.Where("is_system_wallet = ?", true).First(&systemWalEntity).Error
	if err != nil {
		return nil, port.ErrSystemWalletDoesNotExists
	}
	var commissionEntity *types.Commission
	err = r.db.WithContext(ctx).First(&commissionEntity).Error
	if err != nil {
		return nil, err
	}
	commissionee := tr.Amount * commissionEntity.AppCommissionPercentage / 100
	walletTo.Balance += (tr.Amount - commissionee)

	if err := r.db.WithContext(ctx).Save(&walletTo).Error; err != nil {
		return nil, err
	}
	systemWalEntity.Balance += commissionee
	if err := r.db.WithContext(ctx).Save(&systemWalEntity).Error; err != nil {
		return nil, err
	}
	tr.Status = types.TransactionSuccess
	createdTransaction := mapper.DomainTransactionToTransactionEntity(tr)
	err = r.db.WithContext(ctx).Create(&createdTransaction).Error
	if err != nil {
		return nil, err
	}
	return tr, nil
}
