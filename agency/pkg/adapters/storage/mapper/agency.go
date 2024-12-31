package mapper

import (
	"agency/internal/agency/domain"
	"agency/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

func AgencyStorageToDomain(agency types.Agency) *domain.Agency {
	return &domain.Agency{
		ID:        domain.AgencyID(agency.ID),
		CreatedAt: agency.CreatedAt,
		UpdatedAt: agency.UpdatedAt,
		Name:      agency.Name,
		OwnerID:   domain.OwnerID(agency.OwnerID),
	}
}

func AgencyDomainToStorage(agencyDomain domain.Agency) *types.Agency {
	return &types.Agency{
		Model: gorm.Model{
			ID:        uint(agencyDomain.ID),
			CreatedAt: agencyDomain.CreatedAt,
			UpdatedAt: agencyDomain.UpdatedAt,
		},
		Name:    agencyDomain.Name,
		OwnerID: uint(agencyDomain.OwnerID),
	}
}
