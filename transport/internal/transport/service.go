package transport

import (
	"context"
	"log"

	"github.com/onlineTraveling/transport/internal/transport/domain"
	"github.com/onlineTraveling/transport/internal/transport/port"
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}

}

func (s *service) CreateCompanyService(ctx context.Context, vehicle domain.Company) (domain.CompanyID, error) {
	var companyID domain.CompanyID
	companyID, err := s.repo.CreateCompany(ctx, vehicle)
	if err != nil {
		log.Fatalf("Unable to Create Vehicle, error: %v", err)
		return companyID, err
	}
	return companyID, nil
}

func (s *service) UpdateCompanyService(ctx context.Context, vehicle domain.Company) (domain.CompanyID, error) {
	var companyID domain.CompanyID
	companyID, err := s.repo.UpdateCompany(ctx, vehicle)
	if err != nil {
		log.Fatalf("Unable to Update Vehicle, error: %v", err)
		return companyID, err
	}
	return companyID, nil
}

func (s *service) DeleteCompanyService(ctx context.Context, companyID domain.CompanyID) (domain.CompanyID, error) {
	vID, err := s.repo.DeleteCompany(ctx, companyID)
	if err != nil {
		log.Fatalf("Unable to Delete Vehicle, error: %v", err)
		return vID, err
	}
	return vID, nil
}

func (s *service) GetByIDCompanyService(ctx context.Context, companyID domain.CompanyID) (domain.Company, error) {
	vehicle, err := s.repo.GetByIDCompany(ctx, companyID)
	if err != nil {
		log.Fatalf("Unable to Get Vehicle, error: %v", err)
		return domain.Company{}, err
	}
	return vehicle, nil
}

