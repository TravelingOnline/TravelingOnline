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

func (r *vehicleRepo) CreateVehicle(ctx context.Context, vehicle domain.Vehicle) (domain.VehicleID, error) {
	if err:=mapper.VehicleValidator(vehicle); err!=nil{
		return domain.VehicleID(""), err
	}
	// Map the domain.Vehicle to the storage type
	newVehicle := mapper.DomainVehicle2Storage(vehicle)

	// Insert the new vehicle into the database
	if err := r.db.WithContext(ctx).Model(&types.Vehicle{}).Create(&newVehicle).Error; err != nil {
		log.Printf("failed to create vehicle: %v", err)
		return "", fmt.Errorf("unable to create vehicle in the database: %w", err)
	}

	// Return the new vehicle's ID
	return domain.VehicleID(newVehicle.Id), nil
}

func (r *vehicleRepo) UpdateVehicle(ctx context.Context, vehicle domain.Vehicle) (domain.VehicleID, error) {
	if err:=mapper.VehicleValidator(vehicle); err!=nil{
		return domain.VehicleID(""), err
	}
	// Map domain vehicle to storage vehicle model
	updateVehicle := mapper.DomainVehicle2Storage(vehicle)

	// Update the vehicle in the database
	if err := r.db.WithContext(ctx).
		Model(&updateVehicle).
		Where("id = ?", updateVehicle.Id).
		Updates(updateVehicle).Error; err != nil {
		log.Printf("failed to update vehicle with ID %s: %v", updateVehicle.Id, err)
		return domain.VehicleID(""), fmt.Errorf("unable to update vehicle in the database: %w", err)
	}

	// Return the updated vehicle ID
	return domain.VehicleID(updateVehicle.Id), nil
}

func (r *vehicleRepo) DeleteVehicle(ctx context.Context, vehicleID domain.VehicleID) (domain.VehicleID, error) {
	var vID domain.VehicleID
	// Validate input
	if vehicleID == "" {
		return vID, fmt.Errorf("VEHICLE ID CANNOT BE EMPTY")
	}

	// Delete the vehicle from the database
	if err := r.db.WithContext(ctx).
		Where("id = ?", string(vehicleID)).
		Delete(&types.Vehicle{}).Error; err != nil {
		log.Printf("failed to delete vehicle with ID %s: %v", vehicleID, err)
		return vID, fmt.Errorf("unable to delete vehicle from the database: %w", err)
	}

	return vID, nil
}

func (r *vehicleRepo) GetByIDVehicle(ctx context.Context, vehicleID domain.VehicleID) (domain.Vehicle, error) {
	// Validate input
	if vehicleID == "" {
		return domain.Vehicle{}, fmt.Errorf("VEHICLE ID CANNOT BE EMPTY")
	}

	// Initialize a storage model to hold the result
	var storageVehicle types.Vehicle

	// Query the database and preload the Owner data
	if err := r.db.WithContext(ctx).
		Preload("Owner"). // Preload the associated Owner record
		Where("id = ?", string(vehicleID)).
		First(&storageVehicle).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Vehicle{}, fmt.Errorf("vehicle with ID %s not found", string(vehicleID))
		}
		log.Printf("failed to fetch vehicle with ID %s: %v", vehicleID, err)
		return domain.Vehicle{}, fmt.Errorf("unable to fetch vehicle from the database: %w", err)
	}

	// Map the storage model to the domain model
	domainVehicle, err := mapper.VehicleStroage2Domain(storageVehicle)
	if err != nil {
		return domain.Vehicle{}, fmt.Errorf("error in mapper: %w", err)
	}

	return domainVehicle, nil
}

func (r *vehicleRepo) RentVehicle(ctx context.Context, rentReq domain.Vehicle) (domain.Vehicle, error) {
	var bestVehicle types.Vehicle // Use the storage type for querying the database

	// Query the database with filtering and ordering=
	if err := r.db.WithContext(ctx).
		Preload("Owner").                           // Preload the associated Owner record
		Where("passenger >= ?", rentReq.Passenger). // Match vehicles with sufficient passenger capacity
		Where("rent_price <= ?", rentReq.RentPrice).
		Where("type = ?", rentReq.Type).
		Where("is_active = true"). // Vehicle must be active
		Order("passenger ASC").    // Closest matching passenger capacity first
		Order("rent_price ASC").   // Cheapest rent price
		Order("model ASC").        // Oldest model
		Order("created_at ASC").   // Earliest creation date
		First(&bestVehicle).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Vehicle{}, fmt.Errorf("no vehicle found for passenger count %d", rentReq.Passenger)
		}
		log.Printf("Failed to fetch the best vehicle: %v", err)
		return domain.Vehicle{}, fmt.Errorf("unable to fetch vehicle from the database: %w", err)
	}

	// Map the storage model to the domain model
	domainVehicle, err := mapper.VehicleStroage2Domain(bestVehicle)
	if err != nil {
		return domain.Vehicle{}, fmt.Errorf("error in mapper: %w", err)
	}


	return domainVehicle, nil
}
