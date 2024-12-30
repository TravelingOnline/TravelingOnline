package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/onlineTraveling/transport/internal/tour/port"
	"github.com/onlineTraveling/transport/internal/tour/domain"
	"github.com/onlineTraveling/transport/pkg/adapters/storage/mapper"
	"github.com/onlineTraveling/transport/pkg/adapters/storage/types"
	"gorm.io/gorm"
)

type tourRepo struct {
	db *gorm.DB
}

func NewTourRepo(db *gorm.DB) port.Repo {
	return &tourRepo{
		db: db,
	}
}

func (r *tourRepo) CreateTour(ctx context.Context, t domain.Tour) (domain.TourID, error) {
	// Map the domain.Tour to the storage type
	newTour := mapper.DomainTour2Storage(t)

	// Insert the new tour into the database
	if err := r.db.WithContext(ctx).Model(&types.Company{}).Create(&newTour).Error; err != nil {
		log.Printf("failed to create tour: %v", err)
		return "", fmt.Errorf("unable to create tour in the database: %w", err)
	}

	// Return the new Company's ID
	return domain.TourID(newTour.Id), nil
}

func (r *tourRepo) UpdateTour(ctx context.Context, tour domain.Tour) (domain.TourID, error) {
	// Map domain tour to storage tour model
	updateTour := mapper.DomainTour2Storage(tour)

	// Update the tour in the database
	if err := r.db.WithContext(ctx).
		Model(&updateTour).
		Where("id = ?", updateTour.Id).
		Updates(updateTour).Error; err != nil {
		log.Printf("failed to update tour with ID %s: %v", updateTour.Id, err)
		return domain.TourID(""), fmt.Errorf("unable to update tour in the database: %w", err)
	}

	// Return the updated company ID
	return domain.TourID(updateTour.Id), nil
}

func (r *tourRepo) DeleteTour(ctx context.Context, tourID domain.TourID) (domain.TourID, error) {
	var vID domain.TourID
	// Validate input
	if tourID == "" {
		return vID, fmt.Errorf("tour ID cannot be empty")
	}

	// Delete the tour from the database
	if err := r.db.WithContext(ctx).
		Where("id = ?", string(tourID)).
		Delete(&types.Tour{}).Error; err != nil {
		log.Printf("failed to delete tour with ID %s: %v", tourID, err)
		return vID, fmt.Errorf("unable to delete tour from the database: %w", err)
	}

	return vID, nil
}

func (r *tourRepo) GetByIDTour(ctx context.Context, tourID domain.TourID) (domain.Tour, error) {
	// Validate input
	if tourID == "" {
		return domain.Tour{}, fmt.Errorf("tour ID cannot be empty")
	}

	// Initialize a storage model to hold the result
	var storageTour types.Tour

	// Query the database and preload the TechnicalTeamID data
	if err := r.db.WithContext(ctx).
		Preload("TechnicalTeam"). // Preload the associated TechnicalTeamID record
		Where("id = ?", string(tourID)).
		First(&storageTour).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.Tour{}, fmt.Errorf("tour with ID %s not found", tourID)
		}
		log.Printf("failed to fetch tour with ID %s: %v", tourID, err)
		return domain.Tour{}, fmt.Errorf("unable to fetch tour from the database: %w", err)
	}

	// Map the storage model to the domain model
	domainTour, err := mapper.TourStroage2Domain(storageTour)
	if err != nil {
		return domain.Tour{}, fmt.Errorf("error in mapper: %w", err)
	}

	return domainTour, nil
}
