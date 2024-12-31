package vehicle

import (
	"context"
	"log"

	"github.com/onlineTraveling/transport/internal/vehicle/domain"
	"github.com/onlineTraveling/transport/internal/vehicle/port"
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}

}

func (s *service) RentCarService(ctx context.Context, vehicle domain.Vehicle) (domain.VehicleID, error) {
	v, err := s.repo.RentCar(ctx, vehicle)

	if err != nil {
		log.Fatalf("Unable to rent car, error: %v", err)
		return v, err
	}
	return v, nil
}
