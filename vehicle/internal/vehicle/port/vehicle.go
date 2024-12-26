package port

import (
	"context"

	"github.com/onlineTraveling/vehicle/internal/vehicle/domain"
)

type Repo interface {
	CreateVehicle(ctx context.Context, vehicle domain.Vehicle) (domain.VehicleID, error)
}
