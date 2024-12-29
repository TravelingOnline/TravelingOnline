package app

import (
	"user_service/config"

	"gorm.io/gorm"
)

type app struct {
	db  *gorm.DB
	cfg config.Config
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) Config() config.Config {
	return a.cfg
}
