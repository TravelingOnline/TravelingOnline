package agency

import (
	"agency/internal/agency/domain"
	"agency/internal/agency/port"
	"context"
	"errors"
)

var (
	ErrAgencyOnCreate = errors.New("error on creating agency")
	ErrAgencyNotFound = errors.New("couldn't find agency")
	ErrAgencyOnUpdate = errors.New("error on updating agency")
	ErrAgencyOnDelete = errors.New("error on deleting agency")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateAgency(ctx context.Context, agency domain.Agency) error {
	err := s.repo.Create(ctx, agency)

	if err != nil {
		return ErrAgencyOnCreate
	}

	return nil
}

func (s *service) GetAllAgencies(ctx context.Context, page int, pagesize int) ([]domain.Agency, error) {
	agencies, err := s.repo.GetAll(ctx, page, pagesize)

	if err != nil {
		return []domain.Agency{}, ErrAgencyNotFound
	}

	return agencies, nil
}

func (s *service) GetAgencyByID(ctx context.Context, id uint) (*domain.Agency, error) {
	agency, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return &domain.Agency{}, ErrAgencyNotFound
	}

	return agency, nil
}

func (s *service) GetAgencyByOwnerID(ctx context.Context, id uint) (*domain.Agency, error) {
	agency, err := s.repo.GetByOwnerID(ctx, id)
	if err != nil {
		return &domain.Agency{}, ErrAgencyNotFound
	}

	return agency, nil
}

func (s *service) UpdateAgency(ctx context.Context, agency *domain.Agency) error {
	err := s.repo.Update(ctx, agency)

	if err != nil {
		return ErrAgencyOnUpdate
	}

	return nil
}

func (s *service) DeleteAgency(ctx context.Context, id uint) error {
	err := s.repo.Delete(ctx, id)

	if err != nil {
		return ErrAgencyOnDelete
	}

	return nil
}
