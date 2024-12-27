package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/onlineTraveling/transport/internal/transport/domain"
	"github.com/onlineTraveling/transport/internal/transport/port"
	"github.com/onlineTraveling/transport/pkg/adapters/storage/mapper"
	"github.com/onlineTraveling/transport/pkg/adapters/storage/types"
	"gorm.io/gorm"
)

type transportRepo struct {
	db *gorm.DB
}

func NewTransportRepo(db *gorm.DB) port.Repo {
	return &transportRepo{
		db: db,
	}
}

func (r *transportRepo) CreateVehicle(ctx context.Context, v domain.Company) (domain.CompanyID, error) {
	// Map the domain.Vehicle to the storage type
	newVehicle := mapper.DomainCompany2Storage(v)

	// Insert the new vehicle into the database
	if err := r.db.WithContext(ctx).Model(&types.Company{}).Create(&newVehicle).Error; err != nil {
		log.Printf("failed to create company: %v", err)
		return "", fmt.Errorf("unable to create company in the database: %w", err)
	}

	// Return the new vehicle's ID
	return domain.CompanyID(newVehicle.Id), nil
}

func (r *transportRepo) UpdateVehicle(ctx context.Context, vehicle domain.Company) (domain.CompanyID, error) {
	// Map domain vehicle to storage vehicle model
	updateVehicle := mapper.DomainCompany2Storage(vehicle)

	// Update the vehicle in the database
	if err := r.db.WithContext(ctx).
		Model(&updateVehicle).
		Where("id = ?", updateVehicle.Id).
		Updates(updateVehicle).Error; err != nil {
		log.Printf("failed to update vehicle with ID %s: %v", updateVehicle.Id, err)
		return domain.CompanyID(""), fmt.Errorf("unable to update vehicle in the database: %w", err)
	}

	// Return the updated vehicle ID
	return domain.CompanyID(updateVehicle.Id), nil
}

func (r *transportRepo) DeleteVehicle(ctx context.Context, vehicleID domain.CompanyID) (domain.CompanyID, error) {
	var vID domain.CompanyID
	// Validate input
	if vehicleID == "" {
		return vID, fmt.Errorf("vehicle ID cannot be empty")
	}

	// Delete the vehicle from the database
	if err := r.db.WithContext(ctx).
		Where("id = ?", string(vehicleID)).
		Delete(&types.Company{}).Error; err != nil {
		log.Printf("failed to delete vehicle with ID %s: %v", vehicleID, err)
		return vID, fmt.Errorf("unable to delete vehicle from the database: %w", err)
	}

	return vID, nil
}

func (r *transportRepo) GetByIDVehicle(ctx context.Context, vehicleID domain.CompanyID) (domain.Company, error) {
	// Validate input
	if vehicleID == "" {
		return domain.Company{}, fmt.Errorf("vehicle ID cannot be empty")
	}

	// Initialize a storage model to hold the result
	var storageVehicle types.Company

	// Query the database and preload the Owner data
	if err := r.db.WithContext(ctx).
		Preload("Owner"). // Preload the associated Owner record
		Where("id = ?", string(vehicleID)).
		First(&storageVehicle).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Company{}, fmt.Errorf("vehicle with ID %s not found", vehicleID)
		}
		log.Printf("failed to fetch vehicle with ID %s: %v", vehicleID, err)
		return domain.Company{}, fmt.Errorf("unable to fetch vehicle from the database: %w", err)
	}

	// Map the storage model to the domain model
	domainVehicle, err := mapper.CompanyStroage2Domain(storageVehicle)
	if err != nil {
		return domain.Company{}, fmt.Errorf("error in mapper: %w", err)
	}

	return domainVehicle, nil
}
