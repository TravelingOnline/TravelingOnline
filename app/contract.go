package app

import (
	"context"

	"github.com/travelingOnline/config"
	"gorm.io/gorm"
	// userPort "github.com/travelingOnline/internal/user/port"
)

type App interface {
	DB() *gorm.DB

	Config(ctx context.Context) config.Config
	// UserService(ctx context.Context) userPort.Service
}
