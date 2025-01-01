package app

import (
	"context"

	"github.com/onlineTraveling/auth/config"

	"github.com/onlineTraveling/auth/internal/codeVerification"
	"github.com/onlineTraveling/auth/internal/common"
	notifPort "github.com/onlineTraveling/auth/internal/notification/port"

	"github.com/onlineTraveling/auth/internal/user"
	userPort "github.com/onlineTraveling/auth/internal/user/port"
	"github.com/onlineTraveling/auth/pkg/adapters/storage"
	"github.com/onlineTraveling/auth/pkg/postgres"

	codeVerificationPort "github.com/onlineTraveling/auth/internal/codeVerification/port"

	"github.com/go-co-op/gocron/v2"
	"github.com/onlineTraveling/auth/internal/notification"
	appCtx "github.com/onlineTraveling/auth/pkg/context"
	"gorm.io/gorm"
)

type app struct {
	db                *gorm.DB
	cfg               config.Config
	userService       userPort.Service
	notifService      notifPort.Service
	codeVrfctnService codeVerificationPort.Service
}

// CodeVerificationService implements App.

func (a *app) DB() *gorm.DB {
	return a.db
}
func (a *app) UserService(ctx context.Context) userPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.userService == nil {
			a.userService = a.userServiceWithDB(a.db)
		}
		return a.userService
	}

	return a.userServiceWithDB(db)
}

func (a *app) userServiceWithDB(db *gorm.DB) userPort.Service {
	return user.NewService(storage.NewUserRepo(db))
}
func (a *app) NotifService(ctx context.Context) notifPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.notifService == nil {
			a.notifService = a.notifServiceWithDB(a.db)
		}
		return a.notifService
	}

	return a.notifServiceWithDB(db)
}
func (a *app) notifServiceWithDB(db *gorm.DB) notifPort.Service {
	return notification.NewService(storage.NewNotifRepo(db))

}

func (a *app) codeVerificationServiceWithDB(db *gorm.DB) codeVerificationPort.Service {
	return codeVerification.NewService(
		a.userService, storage.NewOutboxRepo(db), storage.NewCodeVerificationRepo(db))
}

func (a *app) CodeVerificationService(ctx context.Context) codeVerificationPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.codeVrfctnService == nil {
			a.codeVrfctnService = a.codeVerificationServiceWithDB(a.db)
		}
		return a.codeVrfctnService
	}

	return a.codeVerificationServiceWithDB(db)
}

func (a *app) Config(ctx context.Context) config.Config {
	return a.cfg
}

func (a *app) setDB() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		User:   a.cfg.DB.User,
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

	if err != nil {
		return err
	}

	return nil
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	a.userService = user.NewService(storage.NewUserRepo(a.db))
	a.codeVrfctnService = codeVerification.NewService(a.userService, storage.NewOutboxRepo(a.db), storage.NewCodeVerificationRepo(a.db))
	a.notifService = notification.NewService(storage.NewNotifRepo(a.db))
	return a, a.registerOutboxHandlers()
}

func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}

func (a *app) registerOutboxHandlers() error {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		return err
	}

	common.RegisterOutboxRunner(a.codeVerificationServiceWithDB(a.db), scheduler)

	scheduler.Start()

	return nil
}
