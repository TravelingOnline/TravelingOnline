package app

import (
	"context"

	"github.com/onlineTraveling/hotel/config"

	"github.com/onlineTraveling/hotel/pkg/postgres"

	"gorm.io/gorm"
)

type app struct {
	db  *gorm.DB
	cfg config.Config
	// userService  userPort.Service
}

func (a *app) DB() *gorm.DB {
	return a.db
}

// func (a *app) UserService(ctx context.Context) userPort.Service {
// 	db := appCtx.GetDB(ctx)
// 	if db == nil {
// 		if a.userService == nil {
// 			a.userService = a.userServiceWithDB(a.db)
// 		}
// 		return a.userService
// 	}

// 	return a.userServiceWithDB(db)
// }

// func (a *app) userServiceWithDB(db *gorm.DB) userPort.Service {
// 	return user.NewService(storage.NewUserRepo(db))
// }

func (a *app) Config(ctx context.Context) config.Config {
	return a.cfg
}

func (a *app) setDB() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		User:   a.cfg.HOTEL_DB.User,
		Pass:   a.cfg.HOTEL_DB.Password,
		Host:   a.cfg.HOTEL_DB.Host,
		Port:   a.cfg.HOTEL_DB.Port,
		DBName: a.cfg.HOTEL_DB.Database,
		Schema: a.cfg.HOTEL_DB.Schema,
	})

	postgres.GormMigrations(db)

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

	// a.userService = user.NewService(storage.NewUserRepo(a.db))
	return a, nil
}

func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}
