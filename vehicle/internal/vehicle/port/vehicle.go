package port

import (
	"context"

	"github.com/onlineTraveling/vehicle/internal/vehicle/domain"
)

type Repo interface {
	CreateVehicle(ctx context.Context, vehicle domain.Vehicle) (domain.VehicleID, error)
	UpdateVehicle(ctx context.Context, vehicle domain.Vehicle) (domain.VehicleID, error)
	DeleteVehicle(ctx context.Context, vehicleID domain.VehicleID) (domain.VehicleID, error)
	GetByIDVehicle(ctx context.Context, vehicleID domain.VehicleID) (domain.Vehicle, error)
	RentVehicle(ctx context.Context, rentReq domain.Vehicle) (domain.Vehicle, error)
}
