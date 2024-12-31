package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/onlineTraveling/transport/config"
	"github.com/onlineTraveling/transport/internal/tour/domain"
	"github.com/onlineTraveling/transport/internal/tour/port"
	bpb "github.com/onlineTraveling/transport/pkg/adapters/storage/bank-pb"
	"github.com/onlineTraveling/transport/pkg/adapters/storage/helpers"
	"github.com/onlineTraveling/transport/pkg/adapters/storage/mapper"
	"github.com/onlineTraveling/transport/pkg/adapters/storage/types"
	vpb "github.com/onlineTraveling/transport/pkg/adapters/storage/vehicle-pb"
	"gorm.io/gorm"
)

type tourRepo struct {
	db     *gorm.DB
	config config.Config
}

func NewTourRepo(db *gorm.DB, config config.Config) port.Repo {
	return &tourRepo{
		db:     db,
		config: config,
	}
}

func (r *tourRepo) CreateTour(ctx context.Context, tour domain.Tour) (domain.TourID, error) {
	if !helpers.ValidDate(tour.StartDate) {
		return domain.TourID(""), fmt.Errorf("Wrong StartDate format")
	}
	if !helpers.ValidDate(tour.EndDate) {
		return domain.TourID(""), fmt.Errorf("Wrong EndDate format")
	}
	// Map the domain.Tour to the storage type
	newTour := mapper.DomainTour2Storage(tour)

	// Insert the new tour into the database
	if err := r.db.WithContext(ctx).Model(&types.Company{}).Create(&newTour).Error; err != nil {
		log.Printf("failed to create tour: %v", err)
		return "", fmt.Errorf("unable to create tour in the database: %w", err)
	}

	// Return the new Company's ID
	return domain.TourID(newTour.Id), nil
}

func (r *tourRepo) UpdateTour(ctx context.Context, tour domain.Tour) (domain.TourID, error) {
	// Validate input
	if tour.Id == "" {
		return domain.TourID(""), fmt.Errorf("TOUR ID CANNOT BE EMPTY")
	}
	if !helpers.ValidDate(tour.StartDate) {
		return domain.TourID(""), fmt.Errorf("WRONG STARTDATE FORMAT")
	}
	if !helpers.ValidDate(tour.EndDate) {
		return domain.TourID(""), fmt.Errorf("WRONG ENDDATE FORMAT")
	}
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
	//Share Money
	if tour.Ended {
		var peoples []string
		senderOwnerID := tour.Company.Owner.Id
		for _, v := range tour.TechnicalTeam {
			peoples = append(peoples, v.Id)
		}
		portion := int(tour.Price)/len(peoples) + 1

		client, conn, err := helpers.NewBankClient(&r.config.Bank.Host, &r.config.Bank.HttpPort)
		if err != nil {
			log.Fatalf("Error creating gRPC client: %v", err)
		}
		defer conn.Close()
		ctx, cancel := context.WithTimeout(ctx, time.Second*10)
		defer cancel()
		for _, v := range peoples {
			req := &bpb.TransferRequest{
				SenderOwnerID:   senderOwnerID,
				ReceiverOwnerID: v,
				Amount:          uint64(portion),
			}
			res, err := client.Transfer(ctx, req)
			if err != nil {
				log.Fatalf("Failed to Transfer money: %v", err)
			}
			log.Printf("Transfer Money: %v", res)
		}
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

func (r *tourRepo) RentCar(ctx context.Context, tType string, passenger int, price int) (domain.Tour, error) {
	client, conn, err := helpers.NewVehicleClient(&r.config.Vehicle.Host, &r.config.Vehicle.HttpPort)
	if err != nil {
		log.Fatalf("Error creating gRPC client: %v", err)
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	rentReq := &vpb.RentVehicleRequest{
		Passenger: int32(passenger),
		Type:      tType,
		Price:     int32(price),
	}
	rentResp, err := client.RentVehicle(ctx, rentReq)
	if err != nil {
		log.Fatalf("Failed to rent vehicle: %v", err)
	}

	log.Printf("Rent Vehicle: %v", rentResp)
	return domain.Tour{
		Vehicle: domain.Vehicle{
			Id:              rentResp.Id,
			Unicode:         rentResp.Unicode,
			RequiredExperts: rentResp.RequiredExperts,
			Speed:           rentResp.Speed,
			RentPrice:       rentResp.RentPrice,
			Type:            rentReq.Type,
			Passenger:       rentReq.Passenger,
			Model:           rentResp.Model,
		},
	}, nil
}
