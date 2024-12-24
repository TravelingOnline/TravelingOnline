package postgres

import (
	"fmt"
	"log"

	"github.com/onlineTraveling/bank/pkg/adapters/storage/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConnOptions struct {
	User   string
	Pass   string
	Host   string
	Port   uint
	DBName string
	Schema string
}

func (o DBConnOptions) PostgresDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		o.Host, o.Port, o.User, o.Pass, o.DBName, o.Schema)
}

func NewPsqlGormConnection(opt DBConnOptions) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(opt.PostgresDSN()), &gorm.Config{
		Logger: logger.Default,
	})
}
func GormMigrations(db *gorm.DB) {
	err := db.AutoMigrate(

		&types.CreditCard{},
		&types.Wallet{},
		&types.WalletCreditCard{},
		&types.BankTransaction{},
		// &types.WalletTransaction{},
		&types.Commission{},
	)
	if err != nil {
		log.Fatalf("failed to migrate models: %v", err)
	}
}
func AddUuidExtension(db *gorm.DB) error {
	return db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error
}
func SeedWalletAndCommisionTableRecords(db *gorm.DB) error {
	err := seedCommissionDB(db)
	if err != nil {
		return err
	}
	err = seedWalletDB(db)
	if err != nil {
		return err
	}
	return nil
}

func seedCommissionDB(db *gorm.DB) error {
	var commission *types.Commission
	err := db.Table("commission").First(&commission).Error

	if err != nil {
		commission = &types.Commission{AppCommissionPercentage: 1}
	}
	err = db.Table("commissions").Create(&commission).Error
	if err != nil {
		return err
	}
	return nil
}
func seedWalletDB(db *gorm.DB) error {
	var systemWallet *types.Wallet
	err := db.Table("wallets").Where("is_system_wallet = ?", true).First(&systemWallet).Error
	if err != nil {
		systemWallet = &types.Wallet{
			IsSystemWallet: true,
			Balance:        0,
		}
		err := db.Create(&systemWallet).Error
		if err != nil {
			return err
		}
	}
	return nil
}
