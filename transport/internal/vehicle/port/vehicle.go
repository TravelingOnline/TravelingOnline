package port

import (
	"context"

	"github.com/onlineTraveling/transport/internal/vehicle/domain"
)

type Repo interface {
	RentCar(ctx context.Context, vehicle domain.Vehicle) (domain.VehicleID, error)
}
