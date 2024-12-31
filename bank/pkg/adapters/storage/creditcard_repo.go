package storage

import (
	"context"
	"errors"
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

func (r *creditCardRepo) DeleteCard(ctx context.Context, creditCardID uuid.UUID) error {
	result := r.db.WithContext(ctx).Table("credit_cards").Where("id = ?", creditCardID).Delete(&types.CreditCard{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("credit card not found")
	}
	return nil
}

func (r *creditCardRepo) UpdateCard(ctx context.Context, creditCard *domain.CreditCard) error {
	// Map domain credit card to database entity
	updatedCard := mapper.CreditCardDomainToEntity(creditCard)

	// Use GORM's Save method to update the record
	result := r.db.WithContext(ctx).Model(&types.CreditCard{}).Where("id = ?", creditCard.ID).Updates(updatedCard)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no records updated, credit card may not exist")
	}
	return nil
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
