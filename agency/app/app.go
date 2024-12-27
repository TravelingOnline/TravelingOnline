package app

import (
	"agency/config"
	"agency/internal/agency"
	agencyPort "agency/internal/agency/port"
	"agency/internal/tour"
	tourPort "agency/internal/tour/port"
	"agency/pkg/adapters/storage"
	appCtx "agency/pkg/context"
	"agency/pkg/postgres"
	"context"

	"gorm.io/gorm"
)

type app struct {
	db            *gorm.DB
	cfg           config.Config
	tourService   tourPort.Service
	agencyService agencyPort.Service
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) TourService(ctx context.Context) tourPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.tourService == nil {
			a.tourService = a.tourServiceWithDB(a.db)
		}
		return a.tourService
	}

	return a.tourServiceWithDB(db)
}

func (a *app) tourServiceWithDB(db *gorm.DB) tourPort.Service {
	return tour.NewService(storage.NewTourRepo(db))
}

func (a *app) AgencyService(ctx context.Context) agencyPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.agencyService == nil {
			a.agencyService = a.agencyServiceWithDB(a.db)
		}
		return a.agencyService
	}

	return a.agencyServiceWithDB(db)
}

func (a *app) agencyServiceWithDB(db *gorm.DB) agencyPort.Service {
	return agency.NewService(storage.NewAgencyRepo(db))
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) setDB() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		User:   a.cfg.AgencyDB.User,
		Pass:   a.cfg.AgencyDB.Password,
		Host:   a.cfg.AgencyDB.Host,
		Port:   a.cfg.AgencyDB.Port,
		DBName: a.cfg.AgencyDB.Database,
		Schema: a.cfg.AgencyDB.Schema,
	})

	if err != nil {
		return err
	}

	a.db = db
	return nil
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	return a, nil
}

func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}
