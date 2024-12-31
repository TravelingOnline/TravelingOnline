package port

import (
	"context"

	"github.com/onlineTraveling/vehicle/internal/vehicle/domain"
)

type Service interface {
	CreateVehicleService(ctx context.Context, vehicle domain.Vehicle) (domain.VehicleID, error)
	UpdateVehicleService(ctx context.Context, vehicle domain.Vehicle) (domain.VehicleID, error)
	DeleteVehicleService(ctx context.Context, vehicleID domain.VehicleID) (domain.VehicleID, error)
	GetByIDVehicleService(ctx context.Context, vehicleID domain.VehicleID) (domain.Vehicle, error)
	RentVehicleService(ctx context.Context, passengerNo int32) (domain.Vehicle, error)
}
