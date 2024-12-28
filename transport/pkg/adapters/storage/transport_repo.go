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

func (r *transportRepo) CreateCompany(ctx context.Context, v domain.Company) (domain.CompanyID, error) {
	// Map the domain.Vehicle to the storage type
	newVehicle := mapper.DomainCompany2Storage(v)

	// Insert the new company into the database
	if err := r.db.WithContext(ctx).Model(&types.Company{}).Create(&newVehicle).Error; err != nil {
		log.Printf("failed to create company: %v", err)
		return "", fmt.Errorf("unable to create company in the database: %w", err)
	}

	// Return the new Company's ID
	return domain.CompanyID(newVehicle.Id), nil
}

func (r *transportRepo) UpdateCompany(ctx context.Context, company domain.Company) (domain.CompanyID, error) {
	// Map domain company to storage company model
	updateVehicle := mapper.DomainCompany2Storage(company)

	// Update the company in the database
	if err := r.db.WithContext(ctx).
		Model(&updateVehicle).
		Where("id = ?", updateVehicle.Id).
		Updates(updateVehicle).Error; err != nil {
		log.Printf("failed to update company with ID %s: %v", updateVehicle.Id, err)
		return domain.CompanyID(""), fmt.Errorf("unable to update company in the database: %w", err)
	}

	// Return the updated company ID
	return domain.CompanyID(updateVehicle.Id), nil
}

func (r *transportRepo) DeleteCompany(ctx context.Context, companyID domain.CompanyID) (domain.CompanyID, error) {
	var vID domain.CompanyID
	// Validate input
	if companyID == "" {
		return vID, fmt.Errorf("company ID cannot be empty")
	}

	// Delete the company from the database
	if err := r.db.WithContext(ctx).
		Where("id = ?", string(companyID)).
		Delete(&types.Company{}).Error; err != nil {
		log.Printf("failed to delete company with ID %s: %v", companyID, err)
		return vID, fmt.Errorf("unable to delete company from the database: %w", err)
	}

	return vID, nil
}

func (r *transportRepo) GetByIDCompany(ctx context.Context, companyID domain.CompanyID) (domain.Company, error) {
	// Validate input
	if companyID == "" {
		return domain.Company{}, fmt.Errorf("company ID cannot be empty")
	}

	// Initialize a storage model to hold the result
	var storageVehicle types.Company

	// Query the database and preload the Owner data
	if err := r.db.WithContext(ctx).
		Preload("Owner"). // Preload the associated Owner record
		Where("id = ?", string(companyID)).
		First(&storageVehicle).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Company{}, fmt.Errorf("company with ID %s not found", companyID)
		}
		log.Printf("failed to fetch company with ID %s: %v", companyID, err)
		return domain.Company{}, fmt.Errorf("unable to fetch company from the database: %w", err)
	}

	// Map the storage model to the domain model
	domainVehicle, err := mapper.CompanyStroage2Domain(storageVehicle)
	if err != nil {
		return domain.Company{}, fmt.Errorf("error in mapper: %w", err)
	}

	return domainVehicle, nil
}
