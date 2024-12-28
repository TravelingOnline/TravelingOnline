package app

import (
	"github.com/onlineTraveling/transport/config"
	"github.com/onlineTraveling/transport/internal/transport"
	"github.com/onlineTraveling/transport/internal/transport/port"
	"github.com/onlineTraveling/transport/pkg/adapters/storage"
	"github.com/onlineTraveling/transport/pkg/postgres"

	"gorm.io/gorm"
)

type App struct {
	db            *gorm.DB
	cfg           config.Config
	transportService port.Service
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

func (a *App) TransportService() port.Service {
	return a.transportService
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
	a.transportService = transport.NewService(storage.NewTransportRepo(a.db))
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
