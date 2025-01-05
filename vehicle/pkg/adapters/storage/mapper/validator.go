package mapper

import (
	"errors"

	"github.com/onlineTraveling/vehicle/internal/vehicle/domain"
)

func VehicleValidator(vehicle domain.Vehicle) error {
	if vehicle.Id == "" {
		return errors.New("vehicle ID is required")
	}

	if vehicle.Unicode == "" {
		return errors.New("vehicle Unicode is required")
	}

	// if vehicle.Owner{
	// 	return errors.New("owner information is missing")
	// }

	if vehicle.Owner.Id == "" || vehicle.Owner.FirstName == "" || vehicle.Owner.LastName == "" || vehicle.Owner.Email == "" {
		return errors.New("owner details are incomplete")
	}

	if vehicle.RequiredExperts <= 0 {
		return errors.New("required experts must be greater than zero")
	}

	if vehicle.Speed <= 0 {
		return errors.New("speed must be greater than zero")
	}

	if vehicle.RentPrice <= 0 {
		return errors.New("rent price must be greater than zero")
	}

	if vehicle.Passenger <= 0 {
		return errors.New("passenger count must be greater than zero")
	}

	if vehicle.Type == "" {
		return errors.New("vehicle type is required")
	}

	return nil
}
