package service

import (
	"agency/api/pb"
	"agency/internal/agency"
	"agency/internal/agency/domain"
	agencyPort "agency/internal/agency/port"
	"context"
)

type AgencyService struct {
	svc agencyPort.Service
}

func NewAgencyService(svc agencyPort.Service) *AgencyService {
	return &AgencyService{svc}
}

var (
	ErrAgencyOnCreate = agency.ErrAgencyOnCreate
	ErrAgencyOnUpdate = agency.ErrAgencyOnUpdate
	ErrAgencyNotFound = agency.ErrAgencyNotFound
	ErrAgencyOnDelete = agency.ErrAgencyOnDelete
)

func (s *AgencyService) CreateAgency(ctx context.Context, req *pb.AgencyCreateRequest) error {
	err := s.svc.CreateAgency(ctx, domain.Agency{
		Name: req.GetName(),
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *AgencyService) GetAgencyByOwnerID(ctx context.Context, id uint) (*domain.Agency, error) {
	agency, err := s.svc.GetAgencyByOwnerID(ctx, id)
	if err != nil {
		return nil, ErrAgencyNotFound
	}

	return agency, nil
}

func (s *AgencyService) GetAgencyByID(ctx context.Context, id uint) (*domain.Agency, error) {
	agency, err := s.svc.GetAgencyByID(ctx, id)
	if err != nil {
		return nil, ErrAgencyNotFound
	}
	return agency, nil

}

func (s *AgencyService) GetAll(ctx context.Context, page, pagesize int) ([]domain.Agency, error) {
	agencies, err := s.svc.GetAllAgencies(ctx, page, pagesize)
	if err != nil {
		return nil, ErrAgencyNotFound
	}

	return agencies, nil
}

// func (s *AgencyService) UpdateAgency(ctx context.Context, req *pb.AgencyUpdateRequest) error {
// 	err := s.svc.UpdateAgency(ctx, &domain.Agency{
// 		ID:   req.GetID(),
		
// 	})
// }

func (s *AgencyService) DeleteAgency(ctx context.Context, id uint) error {
	err := s.svc.DeleteAgency(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
