package service

import (
	"agency/internal/agency"
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

func (s *AgencyService) CreateAgency(ctx context.Context)