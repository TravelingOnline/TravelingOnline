package port

import (
	"context"

	"github.com/onlineTraveling/transport/internal/transport/domain"
)

type Service interface {
	CreateCompanyService(ctx context.Context, company domain.Company) (domain.CompanyID, error)
	UpdateCompanyService(ctx context.Context, company domain.Company) (domain.CompanyID, error)
	DeleteCompanyService(ctx context.Context, companyID domain.CompanyID) (domain.CompanyID, error)
	GetByIDCompanyService(ctx context.Context, companyID domain.CompanyID) (domain.Company, error)
}
