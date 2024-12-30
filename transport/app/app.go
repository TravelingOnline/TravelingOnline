package app

import (
	"github.com/onlineTraveling/transport/config"
	"github.com/onlineTraveling/transport/internal/company"
	"github.com/onlineTraveling/transport/internal/tour"
	companyPort "github.com/onlineTraveling/transport/internal/company/port"
	tourPort "github.com/onlineTraveling/transport/internal/tour/port"
	"github.com/onlineTraveling/transport/pkg/adapters/storage"
	"github.com/onlineTraveling/transport/pkg/postgres"

	"gorm.io/gorm"
)

type App struct {
	db             *gorm.DB
	cfg            config.Config
	companyService companyPort.Service
	tourService    tourPort.Service
}

func (a *App) setDB() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		User:   a.cfg.DB.Username,
		Pass:   a.cfg.DB.Password,
		Host:   a.cfg.DB.Host,
		Port:   a.cfg.DB.Port,
		DBName: a.cfg.DB.Database,
		Schema: a.cfg.DB.Schema,
	})

	postgres.GormMigrations(db)

	if err != nil {
		return err
	}

	a.db = db

	return nil
}

func (a *App) CompanyService() companyPort.Service {
	return a.companyService
}

func (a *App) TourService() tourPort.Service {
	return a.tourService
}

func (a *App) Config() config.Config {
	return a.cfg
}

func NewApp(cfg config.Config) (*App, error) {
	a := &App{
		cfg: cfg,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}
	a.companyService = company.NewService(storage.NewCompanyRepo(a.db))
	a.tourService = tour.NewService(storage.NewTourRepo(a.db))
	return a, nil
}

func NewMustApp(cfg config.Config) *App {
	App, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return App
}

func (a *App) DB() *gorm.DB {
	return a.db
}
