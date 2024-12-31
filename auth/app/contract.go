package app

import (
	"context"

	"github.com/onlineTraveling/auth/config"
	"gorm.io/gorm"
	// userPort "github.com/onlineTraveling/auth/internal/user/port"
)

type App interface {
	DB() *gorm.DB

	Config(ctx context.Context) config.Config
	// UserService(ctx context.Context) userPort.Service
}
