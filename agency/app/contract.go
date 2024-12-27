package app

import (
	"context"

	"agency/config"
	agencyPort "agency/internal/agency/port"
	tourPort "agency/internal/tour/port"

	"gorm.io/gorm"
)

type App interface {
	AgencyService(ctx context.Context) agencyPort.Service
	TourService(ctx context.Context) tourPort.Service
	DB() *gorm.DB
	Config() config.Config
}
