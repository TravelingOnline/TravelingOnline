package mapper

import (
	"errors"

	"github.com/onlineTraveling/vehicle/internal/vehicle/domain"
	"github.com/onlineTraveling/vehicle/pkg/adapters/storage/types"
)

func VehicleStroage2Domain(v types.Vehicle) (domain.Vehicle, error) {
	// Check if the required fields are valid
	if v.Id == "" {
		return domain.Vehicle{}, errors.New("vehicle ID is required")
	}
	if v.Unicode == "" {
		return domain.Vehicle{}, errors.New("vehicle Unicode is required")
	}
	if v.Owner == nil {
		return domain.Vehicle{}, errors.New("owner information is missing")
	}
	if v.Owner.Id == 0 || v.Owner.FirstName == "" || v.Owner.LastName == "" || v.Owner.Email == "" {
		return domain.Vehicle{}, errors.New("owner details are incomplete")
	}

	// Construct the domain vehicle
	vehicle := domain.Vehicle{
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

	// Return the constructed vehicle and nil error if no validation failed
	return vehicle, nil
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
