package service

import (
	"context"
	"fmt"
	"log"

	"github.com/onlineTraveling/vehicle/api/pb"
	"github.com/onlineTraveling/vehicle/internal/vehicle/domain"
	"github.com/onlineTraveling/vehicle/internal/vehicle/port"
)

type VehicleService struct {
	srv port.Service
}

func NewVehicleService(srv port.Service) *VehicleService {
	return &VehicleService{
		srv: srv,
	}
}

func (v *VehicleService) CreateVehicle(ctx context.Context, vehicle *domain.Vehicle) (pb.CreateVehicleResponse, error) {
	// Call the underlying service to create the vehicle
	vehicleID, err := v.srv.CreateVehicle(ctx, *vehicle)
	if err != nil {
		log.Printf("failed to create vehicle: %v", err)
		return pb.CreateVehicleResponse{}, fmt.Errorf("unable to create vehicle: %w", err)
	}

	// Return the response if no error occurred
	return pb.CreateVehicleResponse{
		Id: string(vehicleID),
	}, nil
}
