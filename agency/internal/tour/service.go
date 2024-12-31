package tour

import (
	"agency/internal/tour/domain"
	"agency/internal/tour/port"
	"context"
	"errors"
)

var (
	ErrTourNotFound = errors.New("couldn't find tour")
	ErrTourOnCreate = errors.New("error on creating tour")
	ErrTourOnUpdate = errors.New("error on updating tour")
	ErrTourOnDelete = errors.New("error on deleting tour")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{repo: repo}
}

func (s *service) CreateTour(ctx context.Context, tour domain.Tour) error {
	err := s.repo.Create(ctx, tour)

	if err != nil {
		return ErrTourOnCreate
	}

	return nil
}

func (s *service) GetAllTours(ctx context.Context, agency uint, page int, pagesize int) ([]domain.Tour, error) {
	tours, err := s.repo.GetAll(ctx, agency, page, pagesize)

	if err != nil {
		return []domain.Tour{}, ErrTourNotFound
	}

	return tours, nil
}

func (s *service) GetTourByFilter(ctx context.Context, filter *domain.TourFilter) (*domain.Tour, error) {
	tour, err := s.repo.GetByFilter(ctx, filter)

	if err != nil {
		return &domain.Tour{}, ErrTourNotFound
	}

	return tour, nil
}

func (s *service) GetTourByID(ctx context.Context, id uint) (*domain.Tour, error) {
	tour, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return &domain.Tour{}, ErrTourNotFound
	}

	return tour, nil
}

func (s *service) UpdateTour(ctx context.Context, tour *domain.Tour) error {
	err := s.repo.Update(ctx, tour)

	if err != nil {
		return ErrTourOnUpdate
	}

	return nil
}

func (s *service) DeleteTour(ctx context.Context, id uint) error {
	err := s.repo.Delete(ctx, id)

	if err != nil {
		return ErrTourOnDelete
	}

	return nil
}
