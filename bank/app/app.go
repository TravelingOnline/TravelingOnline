package app

import (
	"context"

	"github.com/onlineTraveling/bank/config"

	"github.com/onlineTraveling/bank/pkg/postgres"

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
		User:   a.cfg.BANK_DB.User,
		Pass:   a.cfg.BANK_DB.Password,
		Host:   a.cfg.BANK_DB.Host,
		Port:   a.cfg.BANK_DB.Port,
		DBName: a.cfg.BANK_DB.Database,
		Schema: a.cfg.BANK_DB.Schema,
	})

	postgres.AddUuidExtension(db)
	postgres.GormMigrations(db)
	postgres.SeedWalletAndCommisionTableRecords(db)
	postgres.SeedWalletAndCommisionTableRecords(db)

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
