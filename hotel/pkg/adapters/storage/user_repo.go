package storage

import (
	"github.com/onlineTraveling/hotel/internal/user/port"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) port.Repo {
	return &userRepo{db}

}
