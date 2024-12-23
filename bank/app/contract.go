package app

import (
	"context"

	"github.com/onlineTraveling/bank/config"
	"gorm.io/gorm"
	// userPort "github.com/onlineTraveling/bank/internal/user/port"
)

type App interface {
	DB() *gorm.DB

	Config(ctx context.Context) config.Config
	// UserService(ctx context.Context) userPort.Service
}
