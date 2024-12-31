package port

import (
	"context"
	"github.com/onlineTraveling/transport/internal/vehicle/domain"
)

type Service interface {
	RentCarService(ctx context.Context, vehicle domain.Vehicle) (domain.VehicleID, error)
}
