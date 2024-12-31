package app

import (
	"context"

	"github.com/onlineTraveling/hotel/config"
	"gorm.io/gorm"
	// userPort "github.com/onlineTraveling/hotel/internal/user/port"
)

type App interface {
	DB() *gorm.DB

	Config(ctx context.Context) config.Config
	// UserService(ctx context.Context) userPort.Service
}
