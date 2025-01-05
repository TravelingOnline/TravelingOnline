package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/onlineTraveling/transport/api/pb"
	"github.com/onlineTraveling/transport/internal/company/domain"
	"github.com/onlineTraveling/transport/internal/company/port"
)

type CompanyService struct {
	srv port.Service
}

func NewCompanyService(srv port.Service) *CompanyService {
	return &CompanyService{
		srv: srv,
	}
}

func (v *CompanyService) CreateCompany(ctx context.Context, req *domain.Company) (*domain.CompanyID, error) {
	req.Id = uuid.New().String()
	cID, err := v.srv.CreateCompanyService(ctx, *req)
	if err != nil {
		return nil, err
	}
	return &cID, nil
}

func (v *CompanyService) UpdateCompany(ctx context.Context, req *domain.Company) (*domain.CompanyID, error) {
	cID, err := v.srv.UpdateCompanyService(ctx, *req)
	if err != nil {
		return nil, err
	}
	return &cID, nil
}

func (v *CompanyService) DeleteCompany(ctx context.Context, vID *domain.CompanyID) (*pb.DeleteCompanyResponse, error) {
	// Call the service to delete the company
	deletedCompanyID, err := v.srv.DeleteCompanyService(ctx, *vID)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteCompanyResponse{
		Id: string(deletedCompanyID),
	}, nil
}

func (v *CompanyService) GetByIDCompany(ctx context.Context, vID *domain.CompanyID) (*domain.Company, error) {
	company, err := v.srv.GetByIDCompanyService(ctx, *vID)
	if err != nil {
		return &domain.Company{}, err
	}
	return &company, nil
}
