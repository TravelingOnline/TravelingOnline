package tour

import (
	"context"
	"fmt"
	"log"

	"github.com/onlineTraveling/transport/internal/tour/domain"
	"github.com/onlineTraveling/transport/internal/tour/port"
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateTourService(ctx context.Context, tour domain.Tour) (domain.TourID, error) {
	var tourID domain.TourID
	v, err := s.repo.RentCar(ctx, tour.Type, int(tour.Capacity), int(tour.Price))
	if err != nil {
		log.Fatalf("Unable to Create Tour, error: %v", err)
		return tourID, err
	}
	tour.Vehicle = v.Vehicle
	tourID, err = s.repo.CreateTour(ctx, tour)
	if err != nil {
		log.Fatalf("Unable to Create Tour, error: %v", err)
		return tourID, err
	}
	return tourID, nil
}

func (s *service) UpdateTourService(ctx context.Context, tour domain.Tour) (domain.TourID, error) {
	var tourID domain.TourID
	oldTour, err := s.repo.GetByIDTour(ctx, domain.TourID(tour.Id))
	if err != nil {
		log.Fatalf("Unable to Update Tour, error: %v", err)
		return tourID, err
	}
	if oldTour.Type != tour.Type || oldTour.Price != tour.Price || oldTour.Capacity != tour.Capacity {
		v, err := s.repo.RentCar(ctx, tour.Type, int(tour.Capacity), int(tour.Price))
		if err != nil {
			log.Fatalf("Unable to Update Tour, error: %v", err)
			return tourID, err
		}
		tour.Vehicle = v.Vehicle
	}

	tourID, err = s.repo.UpdateTour(ctx, tour)
	if err != nil {
		log.Fatalf("Unable to Update Tour, error: %v", err)
		return tourID, err
	}
	return tourID, nil
}

func (s *service) DeleteTourService(ctx context.Context, tourID domain.TourID) (domain.TourID, error) {
	tID, err := s.repo.DeleteTour(ctx, tourID)
	if err != nil {
		log.Fatalf("Unable to Delete Tour, error: %v", err)
		return tID, err
	}
	return tID, nil
}

func (s *service) GetByIDTourService(ctx context.Context, tourID domain.TourID) (domain.Tour, error) {
	tour, err := s.repo.GetByIDTour(ctx, tourID)
	if err != nil {
		log.Fatalf("Unable to Get Tour, error: %v", err)
		return domain.Tour{}, err
	}
	return tour, nil
}
