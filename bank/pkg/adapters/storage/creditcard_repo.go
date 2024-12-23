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

type creditCardRepo struct {
	db *gorm.DB
}

func NewCreditCardRepo(db *gorm.DB) port.CreditCardRepo {
	return &creditCardRepo{
		db: db,
	}
}
func (r *creditCardRepo) CreateCardAndAddToWallet(ctx context.Context, card *domain.CreditCard, userID uuid.UUID) (*domain.CreditCard, error) {
	var userWalletEntity *types.Wallet
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&userWalletEntity).Error; err != nil {
		return nil, err
	}
	newCreditCard := mapper.CreditCardDomainToEntity(card)
	if err := r.db.WithContext(ctx).Where("number = ?", newCreditCard.Number).First(&newCreditCard).Error; err != nil {
		if err = r.db.Create(&newCreditCard).Error; err != nil {
			return nil, err
		}
	}
	walletCreditCardEntity := port.NewWalletCreditCard(userWalletEntity.ID, newCreditCard.ID)
	if err := r.db.WithContext(ctx).Create(&walletCreditCardEntity).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, port.ErrCardAlreadyExists
		}
		return nil, err
	}
	createdCreditCard := mapper.CreditCardEntityToDomain(newCreditCard)
	return createdCreditCard, nil
}

func (r *creditCardRepo) GetUserWalletCards(ctx context.Context, userID uuid.UUID) ([]domain.CreditCard, error) {
	var creditCardEntities []*types.CreditCard

	err := r.db.WithContext(ctx).Joins("JOIN wallet_credit_cards ON wallet_credit_cards.credit_card_id = credit_cards.id").
		Joins("JOIN wallets ON wallets.id = wallet_credit_cards.wallet_id").
		Where("wallets.user_id = ?", userID).
		Find(&creditCardEntities).Error

	if err != nil {
		return nil, err
	}
	allDomainCards := mapper.BatchCreditCardEntityToDomain(creditCardEntities)
	return allDomainCards, nil
}
