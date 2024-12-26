package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/onlineTraveling/vehicle/internal/vehicle/domain"
	"github.com/onlineTraveling/vehicle/internal/vehicle/port"
	"github.com/onlineTraveling/vehicle/pkg/adapters/storage/mapper"
	"github.com/onlineTraveling/vehicle/pkg/adapters/storage/types"
	"gorm.io/gorm"
)

type vehicleRepo struct {
	db *gorm.DB
}

func NewVehicleRepo(db *gorm.DB) port.Repo {
	return &vehicleRepo{
		db: db,
	}
}

func (r *vehicleRepo) CreateVehicle(ctx context.Context, v domain.Vehicle) (domain.VehicleID, error) {
	// Map the domain.Vehicle to the storage type
	newVehicle := mapper.DomainVehicle2Storage(v)

	// Insert the new vehicle into the database
	if err := r.db.WithContext(ctx).Model(&types.Vehicle{}).Create(&newVehicle).Error; err != nil {
		log.Printf("failed to create vehicle: %v", err)
		return "", fmt.Errorf("unable to create vehicle in the database: %w", err)
	}

	// Return the new vehicle's ID
	return domain.VehicleID(newVehicle.Id), nil
}
