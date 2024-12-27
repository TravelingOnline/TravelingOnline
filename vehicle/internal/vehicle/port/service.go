package port

import (
	"context"

	"github.com/onlineTraveling/vehicle/internal/vehicle/domain"
)

type Service interface {
	// CreateVehicleService() (domain.VehicleID, error)
	CreateVehicleService(ctx context.Context, vehicle domain.Vehicle) (domain.VehicleID, error)
}
