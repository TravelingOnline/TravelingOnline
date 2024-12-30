package company

import (
	"context"
	"log"

	"github.com/onlineTraveling/transport/internal/company/domain"
	"github.com/onlineTraveling/transport/internal/company/port"
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}

}

func (s *service) CreateCompanyService(ctx context.Context, company domain.Company) (domain.CompanyID, error) {
	var companyID domain.CompanyID
	companyID, err := s.repo.CreateCompany(ctx, company)
	if err != nil {
		log.Fatalf("Unable to Create Company, error: %v", err)
		return companyID, err
	}
	return companyID, nil
}

func (s *service) UpdateCompanyService(ctx context.Context, company domain.Company) (domain.CompanyID, error) {
	var companyID domain.CompanyID
	companyID, err := s.repo.UpdateCompany(ctx, company)
	if err != nil {
		log.Fatalf("Unable to Update Company, error: %v", err)
		return companyID, err
	}
	return companyID, nil
}

func (s *service) DeleteCompanyService(ctx context.Context, companyID domain.CompanyID) (domain.CompanyID, error) {
	vID, err := s.repo.DeleteCompany(ctx, companyID)
	if err != nil {
		log.Fatalf("Unable to Delete Company, error: %v", err)
		return vID, err
	}
	return vID, nil
}

func (s *service) GetByIDCompanyService(ctx context.Context, companyID domain.CompanyID) (domain.Company, error) {
	company, err := s.repo.GetByIDCompany(ctx, companyID)
	if err != nil {
		log.Fatalf("Unable to Get Company, error: %v", err)
		return domain.Company{}, err
	}
	return company, nil
}

