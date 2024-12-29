package app

import (
	"user_service/config"

	"gorm.io/gorm"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
}
