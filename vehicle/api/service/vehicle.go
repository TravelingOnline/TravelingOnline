package service

import (
	"context"

	"github.com/google/uuid"
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

func (v *VehicleService) CreateVehicle(ctx context.Context, req *domain.Vehicle) (*domain.VehicleID, error) {
	req.Id = uuid.New().String()
	vID, err := v.srv.CreateVehicleService(ctx, *req)
	if err != nil {
		return nil, err
	}
	return &vID, nil
}

func (v *VehicleService) UpdateVehicle(ctx context.Context, req *domain.Vehicle) (*domain.VehicleID, error) {
	vID, err := v.srv.UpdateVehicleService(ctx, *req)
	if err != nil {
		return nil, err
	}
	return &vID, nil
}

func (v *VehicleService) DeleteVehicle(ctx context.Context, vID *domain.VehicleID) (*pb.DeleteVehicleResponse, error) {
	// Call the service to delete the vehicle
	deletedVehcileID, err := v.srv.DeleteVehicleService(ctx, *vID)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteVehicleResponse{
		Id: string(deletedVehcileID),
	}, nil
}

func (v *VehicleService) GetVehicle(ctx context.Context, vID *domain.VehicleID) (*domain.Vehicle, error) {
	vehicle, err := v.srv.GetByIDVehicleService(ctx, *vID)
	if err != nil {
		return &domain.Vehicle{}, err
	}
	return &vehicle, nil
}

func (v *VehicleService) RentVehicle(ctx context.Context, rentReq domain.Vehicle) (*domain.Vehicle, error) {
	bestVehicle, err := v.srv.RentVehicleService(ctx, rentReq)
	if err != nil {
		return &domain.Vehicle{}, err
	}
	return &bestVehicle, nil
}
