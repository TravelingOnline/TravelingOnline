package postgres

import (
	"fmt"
	"log"
	"user_service/pkg/adapters/storage/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConnectionOption struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

func (o DbConnectionOption) PostgresDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		o.Host, o.Port, o.User, o.Password, o.Database,
	)
}

func NewPsqlGormConnection(opt DbConnectionOption) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(opt.PostgresDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&entities.User{}, &entities.Role{}, &entities.Permission{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations completed successfully.")
}
