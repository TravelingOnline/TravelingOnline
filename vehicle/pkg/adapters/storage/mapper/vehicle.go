package mapper

import (
	"github.com/onlineTraveling/vehicle/internal/vehicle/domain"
	"github.com/onlineTraveling/vehicle/pkg/adapters/storage/types"
)

func VehicleStroage2Domain(v types.Vehicle) domain.Vehicle {
	return domain.Vehicle{
		Id:              v.Id,
		Unicode:         v.Unicode,
		RequiredExperts: v.RequiredExperts,
		Speed:           v.Speed,
		RentPrice:       v.RentPrice,
		IsActive:        v.IsActive,
		Type:            v.Type,
		Owner: &domain.Owner{
			Id:        v.Owner.Id,
			FirstName: v.Owner.FirstName,
			LastName:  v.Owner.LastName,
			Email:     v.Owner.Email,
		},
	}
}

func DomainVehicle2Storage(v domain.Vehicle) types.Vehicle {
    return types.Vehicle{
        Id:              v.Id,
        Unicode:         v.Unicode,
        RequiredExperts: v.RequiredExperts,
        Speed:           v.Speed,
        RentPrice:       v.RentPrice,
        IsActive:        v.IsActive,
        Type:            v.Type,
        Owner: &types.Owner{
            Id:        v.Owner.Id,
            FirstName: v.Owner.FirstName,
            LastName:  v.Owner.LastName,
            Email:     v.Owner.Email,
        },
    }
}
