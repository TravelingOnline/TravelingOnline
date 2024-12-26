package vehicle

import (
	"errors"

	"github.com/google/uuid"
	"github.com/onlineTraveling/vehicle/api/pb"
	"github.com/onlineTraveling/vehicle/internal/vehicle/domain"
)

func PBCreateVehicleRequest2DomainVehicle(PBreq *pb.CreateVehicleRequest) (domain.Vehicle, error) {
	// Validate input
	if PBreq == nil {
		return domain.Vehicle{}, errors.New("PBREQ CANNOT BE NIL")
	}
	if PBreq.Owner == nil {
		return domain.Vehicle{}, errors.New("PBREQ.OWNER CANNOT BE NIL")
	}
	if PBreq.Unicode == "" {
		return domain.Vehicle{}, errors.New("UNICODE CANNOT BE EMPTY")
	}
	if PBreq.Owner.Id == 0 || PBreq.Owner.FirstName == "" || PBreq.Owner.LastName == "" || PBreq.Owner.Email == "" {
		return domain.Vehicle{}, errors.New("OWNER DETAILS ARE INCOMPLETE")
	}

	// Construct and return the domain vehicle
	vehicle := domain.Vehicle{
		Id:              uuid.New().String(), // Generate a unique ID
		Unicode:         PBreq.Unicode,
		RequiredExperts: PBreq.RequiredExperts,
		Speed:           PBreq.Speed,
		RentPrice:       PBreq.RentPrice,
		IsActive:        PBreq.IsActive,
		Type:            PBreq.Type,
		Owner: &domain.Owner{
			Id:        PBreq.Owner.Id,
			FirstName: PBreq.Owner.FirstName,
			LastName:  PBreq.Owner.LastName,
			Email:     PBreq.Owner.Email,
		},
	}

	return vehicle, nil
}

func DomainVehicle2PBCreateVehicleRequest(vehicle domain.Vehicle) (*pb.CreateVehicleRequest, error) {
	// Validate input
	if vehicle.Owner == nil {
		return nil, errors.New("vehicle.Owner cannot be nil")
	}

	// Construct and return the protobuf CreateVehicleRequest
	PBreq := &pb.CreateVehicleRequest{
		Unicode:         vehicle.Unicode,
		RequiredExperts: vehicle.RequiredExperts,
		Speed:           vehicle.Speed,
		RentPrice:       vehicle.RentPrice,
		IsActive:        vehicle.IsActive,
		Type:            vehicle.Type,
		Owner: &pb.Owner{
			Id:        vehicle.Owner.Id,
			FirstName: vehicle.Owner.FirstName,
			LastName:  vehicle.Owner.LastName,
			Email:     vehicle.Owner.Email,
		},
	}

	return PBreq, nil
}
