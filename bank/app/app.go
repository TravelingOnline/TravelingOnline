package app

import (
	"context"

	"github.com/onlineTraveling/bank/config"
	"github.com/onlineTraveling/bank/internal/bank"

	"github.com/onlineTraveling/bank/api/service"
	"github.com/onlineTraveling/bank/pkg/adapters/storage"
	"github.com/onlineTraveling/bank/pkg/postgres"
	"github.com/onlineTraveling/bank/pkg/valuecontext"

	"gorm.io/gorm"
)

type App struct {
	db          *gorm.DB
	cfg         config.Config
	bankService *service.BankService
}

func NewApp(cfg config.Config) (*App, error) {
	a := &App{
		cfg: cfg,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}
	a.setBankService()

	return a, nil
}
func (a *App) DB() *gorm.DB {
	return a.db
}

func (a *App) Config(ctx context.Context) config.Config {
	return a.cfg
}

func (a *App) setDB() error {
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

func NewMustApp(cfg config.Config) *App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}
func (a *App) setBankService() {
	if a.bankService != nil {
		return
	}
	a.bankService = service.NewBankService(bank.NewWalletService(storage.NewWalletRepo(a.db)), bank.NewCreditCardService(storage.NewCreditCardRepo(a.db)), bank.NewBankTransactionService(storage.NewBankTransactionRepo(a.db)))
}

func (a *App) BankService() *service.BankService {
	return a.bankService
}

func (a *App) BankServiceFromCtx(ctx context.Context) *service.BankService {
	tx, ok := valuecontext.TryGetTxFromContext(ctx)
	if !ok {
		return a.bankService
	}

	gc, ok := tx.Tx().(*gorm.DB)
	if !ok {
		return a.bankService
	}

	return service.NewBankService(
		bank.NewWalletService(storage.NewWalletRepo(gc)),
		bank.NewCreditCardService(storage.NewCreditCardRepo(gc)),
		bank.NewBankTransactionService(storage.NewBankTransactionRepo(gc)))
}
