package storage

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/onlineTraveling/bank/internal/bank/domain"
	"github.com/onlineTraveling/bank/internal/bank/port"
	"github.com/onlineTraveling/bank/pkg/adapters/storage/mapper"
	"github.com/onlineTraveling/bank/pkg/adapters/storage/types"
	"gorm.io/gorm"
)

type walletRepo struct {
	db *gorm.DB
}

func NewWalletRepo(db *gorm.DB) port.WalletRepo {
	return &walletRepo{
		db: db,
	}
}
func (r *walletRepo) Create(ctx context.Context, wl *domain.Wallet) (*domain.Wallet, error) {
	newWallet := mapper.WalletDomainToEntity(wl)
	err := r.db.WithContext(ctx).Create(&newWallet).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, port.ErrUserAlreadyHasWallet
		}
		return nil, err
	}
	createdWallet := mapper.WalletEntityToDomain(newWallet)
	return createdWallet, nil
}

func (r *walletRepo) Deposit(ctx context.Context, card *domain.CreditCard, amount uint, userID uuid.UUID) (*domain.Wallet, error) {
	var userWalletEntity *types.Wallet
	var cardEntity *types.CreditCard

	if err := r.db.Where("user_id = ?", userID).First(&userWalletEntity).Error; err != nil {
		return nil, err
	}

	// Check if the credit card exists and belongs to the user's wallet
	if err := r.db.WithContext(ctx).Joins("JOIN wallet_credit_cards ON wallet_credit_cards.credit_card_id = credit_cards.id").
		Joins("JOIN wallets ON wallets.id = wallet_credit_cards.wallet_id").
		Where("credit_cards.number = ? AND wallets.user_id = ?", card.Number, userID).
		First(&cardEntity).Error; err != nil {
		return nil, err
	}

	// Increase the wallet balance
	userWalletEntity.Balance += amount
	if err := r.db.WithContext(ctx).Save(&userWalletEntity).Error; err != nil {
		return nil, err
	}

	createdWallet := mapper.WalletEntityToDomain(userWalletEntity)
	return createdWallet, nil
}

func (r *walletRepo) Withdraw(ctx context.Context, card *domain.CreditCard, amount uint, userID uuid.UUID) (*domain.Wallet, error) {
	var userWalletEntity *types.Wallet
	var cardEntity *types.CreditCard

	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&userWalletEntity).Error; err != nil {
		return nil, err
	}

	// Check if the credit card exists and belongs to the user's wallet
	if err := r.db.WithContext(ctx).Joins("JOIN wallet_credit_cards ON wallet_credit_cards.credit_card_id = credit_cards.id").
		Joins("JOIN wallets ON wallets.id = wallet_credit_cards.wallet_id").
		Where("credit_cards.number = ? AND wallets.user_id = ?", card.Number, userID).
		First(&cardEntity).Error; err != nil {
		return nil, err
	}

	if userWalletEntity.Balance < amount {
		return nil, port.ErrNotEnoughBalance
	}
	userWalletEntity.Balance -= amount
	if err := r.db.WithContext(ctx).Save(&userWalletEntity).Error; err != nil {
		return nil, err
	}
	createdWallet := mapper.WalletEntityToDomain(userWalletEntity)
	return createdWallet, nil
}

func (r *walletRepo) GetWallet(ctx context.Context, userID uuid.UUID) (*domain.Wallet, error) {
	var userWalletEntity *types.Wallet
	err := r.db.Where("user_id = ?", userID).First(&userWalletEntity).Error
	if err != nil {
		return nil, err
	}
	var fetchedWalletEntity *types.Wallet
	err = r.db.WithContext(ctx).Model(&types.Wallet{}).Where("id = ?", userWalletEntity.ID).First(&fetchedWalletEntity).Error
	if err != nil {
		return nil, err
	}
	fetchedWalletDomain := mapper.WalletEntityToDomain(fetchedWalletEntity)
	return fetchedWalletDomain, nil
}

func (r *walletRepo) Transfer(ctx context.Context, tr *domain.BankTransaction) (*domain.BankTransaction, error) {
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
