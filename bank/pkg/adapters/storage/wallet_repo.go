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
func (r *walletRepo) DeleteWallet(ctx context.Context, userID uuid.UUID) error {
	var walletEntity *types.Wallet
	err := r.db.Where("user_id =?", userID).First(&walletEntity).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(walletEntity).Error
	return err
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
