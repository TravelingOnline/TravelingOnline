package service

import (
	"context"

	"github.com/onlineTraveling/vehicle/internal/vehicle"
	"github.com/onlineTraveling/vehicle/internal/vehicle/domain"
)

type VehicleService struct {
	srv *vehicle.Service
}

func NewVehicleService(srv *vehicle.Service) *VehicleService {
	return &VehicleService{

		srv: srv,
	}
}

func (v *VehicleService) CreateVehicle(ctx context.Context, req *domain.Vehicle) (*domain.VehicleID, error) {
	// Call the underlying service to create the vehicle
	// TODO mapping req to domain.Vehicle
	// newVehicle := domain.Vehicle{
	// 	Id:              uuid.New().String(), // Generate a unique ID
	// 	Unicode:         req.Unicode,
	// 	RequiredExperts: req.RequiredExperts,
	// 	Speed:           req.Speed,
	// 	RentPrice:       req.RentPrice,
	// 	IsActive:        req.IsActive,
	// 	Type:            req.Type,
	// 	Owner: &domain.Owner{
	// 		Id:        req.Owner.Id,
	// 		FirstName: req.Owner.FirstName,
	// 		LastName:  req.Owner.LastName,
	// 		Email:     req.Owner.Email,
	// 	},
	// }
	vID, err := v.CreateVehicle(ctx, req)
	if err != nil {
		return nil, err
	}
	return vID, nil
	// vehicleID, err := v.srv.CreateVehicle(ctx, newVehicle)
	// if err != nil {
	// 	log.Printf("failed to create vehicle: %v", err)
	// 	return &pb.CreateVehicleResponse{}, fmt.Errorf("unable to create vehicle: %w", err)
	// }

	// // Return the response if no error occurred
	// return &pb.CreateVehicleResponse{
	// 	Id: string(vehicleID),
	// }, nil
}
