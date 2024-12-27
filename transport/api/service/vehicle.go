package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/onlineTraveling/transport/api/pb"
	"github.com/onlineTraveling/transport/internal/transport/domain"
	"github.com/onlineTraveling/transport/internal/transport/port"
)

type TransportService struct {
	srv port.Service
}

func NewTransportService(srv port.Service) *TransportService {
	return &TransportService{
		srv: srv,
	}
}

func (v *TransportService) CreateCompany(ctx context.Context, req *domain.Company) (*domain.CompanyID, error) {
	req.Id = uuid.New().String()
	vID, err := v.srv.CreateCompanyService(ctx, *req)
	if err != nil {
		return nil, err
	}
	return &vID, nil
}

func (v *TransportService) UpdateCompany(ctx context.Context, req *domain.Company) (*domain.CompanyID, error) {
	cID, err := v.srv.UpdateCompanyService(ctx, *req)
	if err != nil {
		return nil, err
	}
	return &cID, nil
}

func (v *TransportService) DeleteCompany(ctx context.Context, vID *domain.CompanyID) (*pb.DeleteCompanyResponse, error) {
	// Call the service to delete the company
	deletedCompanyID, err := v.srv.DeleteCompanyService(ctx, *vID)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteCompanyResponse{
		Id: string(deletedCompanyID),
	}, nil
}

func (v *TransportService) GetByIDCompany(ctx context.Context, vID *domain.CompanyID) (*domain.Company, error) {
	company, err := v.srv.GetByIDCompanyService(ctx, *vID)
	if err != nil {
		return &domain.Company{}, err
	}
	return &company, nil
}
