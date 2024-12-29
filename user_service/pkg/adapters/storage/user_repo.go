package storage

import (
	"context"
	"errors"
	"strings"
	"user_service/internal/user/domain"
	"user_service/internal/user/port"
	"user_service/pkg/adapters/storage/entities"
	"user_service/pkg/adapters/storage/mapper"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) port.Repo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, userDomain domain.User) (domain.UserId, error) {
	newUser := mapper.MapInternalUserToEntity(userDomain)
	err := r.db.WithContext(ctx).Create(&newUser).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return 0, err
		}
		return 0, err
	}
	createdUser := mapper.MapEntityUserToInternal(newUser)
	return createdUser.ID, nil

}
func (r *userRepo) GetById(ctx context.Context, userId domain.UserId) (*domain.User, error) {
	var u entities.User

	err := r.db.WithContext(ctx).Model(&entities.User{}).Where("id = ?", userId).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	User := mapper.MapEntityUserToInternal(u)

	return &User, nil
}
