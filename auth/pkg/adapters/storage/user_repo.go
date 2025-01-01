package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/onlineTraveling/auth/internal/user/domain"
	"github.com/onlineTraveling/auth/internal/user/port"
	"github.com/onlineTraveling/auth/pkg/adapters/storage/mapper"
	"github.com/onlineTraveling/auth/pkg/adapters/storage/types"
	"github.com/onlineTraveling/auth/pkg/helper"
	"github.com/onlineTraveling/auth/pkg/jwt"
	"github.com/onlineTraveling/auth/pkg/logger"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

// GetUserIDByToken implements port.Repo.

func NewUserRepo(db *gorm.DB) port.Repo {
	return &userRepo{db}

}

func (r *userRepo) CreateUser(ctx context.Context, userDomain domain.User) (domain.UserID, error) {
	user := mapper.UserDomain2Storage(userDomain)
	er := r.db.Table("users").WithContext(ctx).Create(user).Error
	if er != nil {
		return domain.UserID(uuid.Nil), er
	}
	er = helper.CreateWalletGrpc(ctx, user.ID)
	if er != nil {
		return domain.UserID(user.ID), er
	}

	return domain.UserID(user.ID), nil
}

func (r *userRepo) GetUserByID(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	var user types.User
	err := r.db.Debug().Table("users").
		Where("id = ?", uuid.UUID(userID)).WithContext(ctx).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	if user.ID == uuid.Nil {
		return nil, nil
	}

	return mapper.UserStorage2Domain(user), nil
}
func (r *userRepo) GetUserIDByToken(ctx context.Context, token string) (*jwt.UserClaims, error) {

	var user types.User
	secret := "ah3*&891809^%$$@$EGJNnjhjkh876$%#@#%"
	claims, err := jwt.ParseToken(token, []byte(secret))
	if err != nil {
		log.Println("Error:", err)
		return nil, err
	}

	err = r.db.Debug().Table("users").
		Where("id = ?", uuid.UUID(claims.UserID)).WithContext(ctx).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	if user.ID == uuid.Nil {
		return nil, nil
	}

	return claims, nil
}
func (r *userRepo) GetUserByEmail(ctx context.Context, email domain.Email) (*domain.User, error) {
	var user types.User
	err := r.db.Table("users").
		Where("email = ?", email).
		First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if user.ID == uuid.Nil {
		return nil, nil
	}

	return mapper.UserStorage2Domain(user), nil
}

func (r *userRepo) GetUserByFilter(ctx context.Context, filter *domain.UserFilter) (*domain.User, error) {
	var user types.User

	q := r.db.Table("users").Debug().WithContext(ctx)

	if filter.ID != domain.UserID(uuid.Nil) {
		q = q.Where("id = ?", uuid.UUID(filter.ID))
	}

	if len(filter.Phone) > 0 {
		q = q.Where("phone = ?", filter.Phone)
	}

	err := q.First(&user).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if user.ID == uuid.Nil {
		return nil, nil
	}

	return mapper.UserStorage2Domain(user), nil
}

func (r *userRepo) UpdateUser(ctx context.Context, user domain.User) error {
	var preUpdateUser types.User
	fmt.Printf("111111,    %v\n     %v",uuid.UUID(user.ID),user.ID)
	err := r.db.Model(&types.User{}).Where("id = ?", uuid.UUID(user.ID)).First((&preUpdateUser)).Error
	if err != nil {
		logger.Error(err.Error(), nil)
		return err
	}
	currentTime := time.Now()
	if currentTime.Sub(preUpdateUser.CreatedAt) > 24*time.Hour {
		return errors.New("can not update user due to limitation of update time")
	}
	updates := make(map[string]interface{})

	if user.Email != "" {
		updates["email"] = user.Email
	}

	tx := r.db.Begin()
	if tx.Error != nil {
		logger.Error(tx.Error.Error(), nil)
		return tx.Error
	}

	// Update the user record
	if err := tx.Model(&types.User{}).Where("id = ?", uuid.UUID(user.ID)).Updates(updates).Error; err != nil {
		logger.Error(err.Error(), nil)
		tx.Rollback()
		return err
	}

	// Commit the transaction
	return tx.Commit().Error
}

func (r *userRepo) DeleteByID(ctx context.Context, userID domain.UserID) error {
	return r.db.Where("id = ?", uuid.UUID(userID)).Delete(&types.User{}).Error
}
