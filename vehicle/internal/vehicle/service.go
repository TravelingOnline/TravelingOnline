package vehicle

import (
	"context"
	"log"

	"github.com/onlineTraveling/vehicle/internal/vehicle/domain"
	"github.com/onlineTraveling/vehicle/internal/vehicle/port"
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}

}

func (s *service) CreateVehicleService(ctx context.Context, vehicle domain.Vehicle) (domain.VehicleID, error) {
	var vehicleID domain.VehicleID
	vehicleID, err := s.repo.CreateVehicle(ctx, vehicle)
	if err != nil {
		log.Fatalf("Unable to Create Vehicle, error: %v", err)
		return vehicleID, err
	}
	return vehicleID, nil
}

func (s *service) UpdateVehicleService(ctx context.Context, vehicle domain.Vehicle) (domain.VehicleID, error) {
	var vehicleID domain.VehicleID
	vehicleID, err := s.repo.UpdateVehicle(ctx, vehicle)
	if err != nil {
		log.Fatalf("Unable to Update Vehicle, error: %v", err)
		return vehicleID, err
	}
	return vehicleID, nil
}

func (s *service) DeleteVehicleService(ctx context.Context, vehicleID domain.VehicleID) (domain.VehicleID, error) {
	vID, err := s.repo.DeleteVehicle(ctx, vehicleID)
	if err != nil {
		log.Fatalf("Unable to Delete Vehicle, error: %v", err)
		return vID, err
	}
	return vID, nil
}

func (s *service) GetByIDVehicleService(ctx context.Context, vehicleID domain.VehicleID) (domain.Vehicle, error) {
	vehicle, err := s.repo.GetByIDVehicle(ctx, vehicleID)
	if err != nil {
		log.Fatalf("Unable to Get Vehicle, error: %v", err)
		return domain.Vehicle{}, err
	}
	return vehicle, nil
}

func (s *service) RentVehicleService(ctx context.Context, passengerNo int32) (domain.Vehicle, error) {
	bestVehicle, err := s.repo.RentVehicle(ctx, passengerNo)
	if err!=nil{
		log.Fatalf("Unable to Get Best Vehicle, error: %v", err)
		return domain.Vehicle{}, err
	}
	return bestVehicle, nil
}
