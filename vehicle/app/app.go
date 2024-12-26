package app

import (
	"log"

	"github.com/onlineTraveling/vehicle/api/service"
	"github.com/onlineTraveling/vehicle/config"
	"github.com/onlineTraveling/vehicle/internal/vehicle"
	"github.com/onlineTraveling/vehicle/pkg/adapters/storage"
	"github.com/onlineTraveling/vehicle/pkg/postgres"

	"gorm.io/gorm"
)

type App struct {
	db             *gorm.DB
	cfg            config.Config
	vehicleService *service.VehicleService
}

func NewApp(cfg config.Config) (*App, error) {
	a := &App{
		cfg: cfg,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}
	a.setVehicleService()

	return a, nil
}

func (a *App) DB() *gorm.DB {
	return a.db
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

	if err != nil {
		return err
	}

	postgres.GormMigrations(db)

	a.db = db

	return nil
}

func (a *App) Config() config.Config {
	return a.cfg
}

func (a *App) VehicleService() service.VehicleService {
	return a.VehicleService()
}

func (a *App) setVehicleService() {
	if a.vehicleService != nil {
		return
	}
	a.vehicleService = service.NewVehicleService(vehicle.NewService(storage.NewVehicleRepo(a.db)))
}

func NewMustApp(cfg config.Config) *App {
	App, err := NewApp(cfg)
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	return App
}
