package postgres

import (
	"fmt"

	"github.com/onlineTraveling/vehicle/pkg/adapters/storage/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	return gorm.Open(postgres.Open(opt.PostgresDSN()))
}
func GormMigrations(db *gorm.DB) {

	db.AutoMigrate(
		&types.Vehicle{},
	)
}
