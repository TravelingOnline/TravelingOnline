package app

import (
	"context"
	"log"

	"github.com/onlineTraveling/bank/config"
	"github.com/onlineTraveling/bank/internal/bank"

	"github.com/onlineTraveling/bank/api/service"

	"github.com/onlineTraveling/bank/pkg/adapters/consul"
	"github.com/onlineTraveling/bank/pkg/adapters/rabbitmq"
	"github.com/onlineTraveling/bank/pkg/adapters/storage"
	"github.com/onlineTraveling/bank/pkg/ports"

	"github.com/onlineTraveling/bank/pkg/adapters/clients/grpc"
	"github.com/onlineTraveling/bank/pkg/ports/clients/clients"
	"github.com/onlineTraveling/bank/pkg/postgres"
	"github.com/onlineTraveling/bank/pkg/valuecontext"

	"gorm.io/gorm"
)

type App struct {
	db              *gorm.DB
	cfg             config.Config
	bankService     *service.BankService
	messageBroker   ports.IMessageBroker
	serviceRegistry ports.IServiceRegistry
	authClient      clients.IAuthClient
}

func NewApp(cfg config.Config) (*App, error) {
	a := &App{
		cfg: cfg,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}
	a.mustRegisterService()
	a.setAuthClient(cfg.Server.ServiceRegistry.AuthServiceName)
	a.setMessageBroker()
	a.setBankService()

	return a, nil
}
func (a *App) DB() *gorm.DB {
	return a.db
}
func (a *App) GetConfig() config.Config {
	return a.cfg
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
func (a *App) MessageBroker() ports.IMessageBroker {
	return a.messageBroker
}

func (a *App) setMessageBroker() {
	messageBrokerCfg := a.cfg.MessageBroker
	if a.messageBroker != nil {
		return
	}
	a.messageBroker = rabbitmq.NewRabbitMQ(messageBrokerCfg.Username, messageBrokerCfg.Password, messageBrokerCfg.Host, messageBrokerCfg.Port)
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
func (a *App) RawDBConnection() *gorm.DB {
	return a.db
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

func (a *App) mustRegisterService() {
	srvCfg := a.cfg.Server
	registry := consul.NewConsul(srvCfg.ServiceRegistry.Address)
	// println("ooooooooooooo ServiceName", srvCfg.ServiceRegistry.ServiceName)
	// println("ooooooooooooo ServiceHostAddress", srvCfg.ServiceHostAddress)
	// println("ooooooooooooo ServiceHTTPPrefixPath", srvCfg.ServiceHTTPPrefixPath)
	// println("ooooooooooooo ServiceHTTPHealthPath", srvCfg.ServiceHTTPHealthPath)
	// println("ooooooooooooo GRPCPort HttpPort", srvCfg.GRPCPort, srvCfg.HttpPort)
	err := registry.RegisterService(srvCfg.ServiceRegistry.ServiceName, srvCfg.ServiceHostAddress, srvCfg.ServiceHTTPPrefixPath, srvCfg.ServiceHTTPHealthPath, srvCfg.GRPCPort, srvCfg.HttpPort)
	if err != nil {
		log.Fatalf("Failed to register service with Consul: %v", err)
	}
	a.serviceRegistry = registry
}

func (a *App) AuthClient() clients.IAuthClient {
	return a.authClient
}

func (a *App) setAuthClient(authServiceName string) {
	if a.authClient != nil {
		return
	}
	a.authClient = grpc.NewGRPCAuthClient(a.serviceRegistry, authServiceName)
}
