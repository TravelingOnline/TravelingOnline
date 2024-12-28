package mapper

import (
	"errors"

	"github.com/onlineTraveling/transport/internal/transport/domain"
	"github.com/onlineTraveling/transport/pkg/adapters/storage/types"
)

func CompanyStroage2Domain(v types.Company) (domain.Company, error) {
	// Check if the required fields are valid
	if v.Id == "" {
		return domain.Company{}, errors.New("company ID is required")
	}
	if v.Owner == nil {
		return domain.Company{}, errors.New("owner information is missing")
	}
	if v.Owner.Id == 0 || v.Owner.FirstName == "" || v.Owner.LastName == "" || v.Owner.Email == "" {
		return domain.Company{}, errors.New("owner details are incomplete")
	}

	// Construct the domain company
	vehicle := domain.Company{
		Id:   v.Id,
		Name: v.Name,
		Owner: &domain.Owner{
			Id:        v.Owner.Id,
			FirstName: v.Owner.FirstName,
			LastName:  v.Owner.LastName,
			Email:     v.Owner.Email,
		},
	}

	// Return the constructed company and nil error if no validation failed
	return vehicle, nil
}

func DomainCompany2Storage(v domain.Company) types.Company {
	return types.Company{
		Id:   v.Id,
		Name: v.Name,
		Owner: &types.Owner{
			Id:        v.Owner.Id,
			FirstName: v.Owner.FirstName,
			LastName:  v.Owner.LastName,
			Email:     v.Owner.Email,
		},
	}
}
