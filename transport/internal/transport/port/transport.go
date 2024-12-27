package port

import (
	"context"

	"github.com/onlineTraveling/transport/internal/transport/domain"
)

type Repo interface {
	CreateCompany(ctx context.Context, company domain.Company) (domain.CompanyID, error)
	UpdateCompany(ctx context.Context, company domain.Company) (domain.CompanyID, error)
	DeleteCompany(ctx context.Context, companyID domain.CompanyID) (domain.CompanyID, error)
	GetByIDCompany(ctx context.Context, companyID domain.CompanyID) (domain.Company, error)
}
