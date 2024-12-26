package vehicle

import (
	"context"
	"log"

	"github.com/onlineTraveling/vehicle/internal/vehicle/domain"
	"github.com/onlineTraveling/vehicle/internal/vehicle/port"
)

type Service struct {
	repo port.Repo
}

func NewService(repo port.Repo) *Service {
	return &Service {
		repo: repo,
	}

}

func (s *Service) CreateVehicle(ctx context.Context, vehicle domain.Vehicle) (domain.VehicleID, error){
	var vehicleID domain.VehicleID
	vehicleID, err := s.repo.CreateVehicle(ctx, vehicle)
	if err!=nil{
		log.Fatalf("Unable to Create Vehicle, error: %v", err)
		return vehicleID, err
	}
	return vehicleID, nil
} 