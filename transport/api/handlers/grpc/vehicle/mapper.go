package transport

import (
	"errors"

	"github.com/onlineTraveling/transport/api/pb"
	"github.com/onlineTraveling/transport/internal/transport/domain"
)

func PBCompanyRequest2DomainCompany(request interface{}) (domain.Company, error) {
	var PBreq struct {
		Id              string
		Unicode         string
		RequiredExperts int32
		Speed           int32
		RentPrice       int32
		IsActive        bool
		Type            string
		Passenger       int
		Model           int
		Owner           *struct {
			Id        uint64
			FirstName string
			LastName  string
			Email     string
		}
	}

	// Map the input to the common structure
	switch req := request.(type) {
	case *pb.CreateCompanyRequest:
		PBreq = struct {
			Id              string
			Unicode         string
			RequiredExperts int32
			Speed           int32
			RentPrice       int32
			IsActive        bool
			Type            string
			Passenger       int
			Model           int
			Owner           *struct {
				Id        uint64
				FirstName string
				LastName  string
				Email     string
			}
		}{
			Id:              req.Id,
			Unicode:         req.Unicode,
			RequiredExperts: req.RequiredExperts,
			Speed:           req.Speed,
			RentPrice:       req.RentPrice,
			IsActive:        req.IsActive,
			Type:            req.Type,
			Passenger:       int(req.Passenger),
			Model:           int(req.Model),
			Owner: &struct {
				Id        uint64
				FirstName string
				LastName  string
				Email     string
			}{
				Id:        req.Owner.Id,
				FirstName: req.Owner.FirstName,
				LastName:  req.Owner.LastName,
				Email:     req.Owner.Email,
			},
		}
	case *pb.UpdateCompanyRequest:
		PBreq = struct {
			Id              string
			Unicode         string
			RequiredExperts int32
			Speed           int32
			RentPrice       int32
			IsActive        bool
			Type            string
			Passenger       int
			Model           int
			Owner           *struct {
				Id        uint64
				FirstName string
				LastName  string
				Email     string
			}
		}{
			Id:              req.Id,
			Unicode:         req.Unicode,
			RequiredExperts: req.RequiredExperts,
			Speed:           req.Speed,
			RentPrice:       req.RentPrice,
			IsActive:        req.IsActive,
			Type:            req.Type,
			Passenger:       int(req.Passenger),
			Model:           int(req.Model),
			Owner: &struct {
				Id        uint64
				FirstName string
				LastName  string
				Email     string
			}{
				Id:        req.Owner.Id,
				FirstName: req.Owner.FirstName,
				LastName:  req.Owner.LastName,
				Email:     req.Owner.Email,
			},
		}
	default:
		return domain.Company{}, errors.New("unsupported request type")
	}

	// Validate input
	if PBreq.Owner == nil {
		return domain.Company{}, errors.New("OWNER CANNOT BE NIL")
	}
	if PBreq.Unicode == "" {
		return domain.Company{}, errors.New("UNICODE CANNOT BE EMPTY")
	}
	if PBreq.Owner.Id == 0 || PBreq.Owner.FirstName == "" || PBreq.Owner.LastName == "" || PBreq.Owner.Email == "" {
		return domain.Company{}, errors.New("OWNER DETAILS ARE INCOMPLETE")
	}

	// Construct and return the domain vehicle
	vehicle := domain.Company{
		Id:              PBreq.Id,
		Unicode:         PBreq.Unicode,
		RequiredExperts: PBreq.RequiredExperts,
		Speed:           PBreq.Speed,
		RentPrice:       PBreq.RentPrice,
		IsActive:        PBreq.IsActive,
		Type:            PBreq.Type,
		Owner:           &domain.Owner{Id: PBreq.Owner.Id, FirstName: PBreq.Owner.FirstName, LastName: PBreq.Owner.LastName, Email: PBreq.Owner.Email},
		Passenger:       PBreq.Passenger,
		Model:           PBreq.Model,
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
		Id:              "",
		Unicode:         vehicle.Unicode,
		RequiredExperts: vehicle.RequiredExperts,
		Speed:           vehicle.Speed,
		RentPrice:       vehicle.RentPrice,
		IsActive:        vehicle.IsActive,
		Type:            vehicle.Type,
		Owner:           &pb.Owner{Id: vehicle.Owner.Id, FirstName: vehicle.Owner.FirstName, LastName: vehicle.Owner.LastName, Email: vehicle.Owner.Email},
		Passenger:       int32(vehicle.Passenger),
		Model:           int32(vehicle.Model),
	}

	return PBreq, nil
}

func DomainVehicle2PBVehicleResponse(vehicle domain.Vehicle) (*pb.GetVehicleResponse, error) {
	// Validate input
	if vehicle.Owner == nil {
		return nil, errors.New("vehicle.Owner cannot be nil")
	}

	// Construct and return the protobuf CreateVehicleRequest
	PBreq := &pb.GetVehicleResponse{
		Unicode:         vehicle.Unicode,
		RequiredExperts: vehicle.RequiredExperts,
		Speed:           vehicle.Speed,
		RentPrice:       vehicle.RentPrice,
		IsActive:        vehicle.IsActive,
		Type:            vehicle.Type,
		Passenger:       int32(vehicle.Passenger),
		Model:           int32(vehicle.Model),

		Owner: &pb.Owner{
			Id:        vehicle.Owner.Id,
			FirstName: vehicle.Owner.FirstName,
			LastName:  vehicle.Owner.LastName,
			Email:     vehicle.Owner.Email,
		},
	}

	return PBreq, nil
}

func DomainVehicle2PBRentVehicleRequest(vehicle domain.Vehicle) (*pb.RentVehicleResponse, error) {
	// Validate input
	if vehicle.Owner == nil {
		return nil, errors.New("vehicle.Owner cannot be nil")
	}

	// Construct and return the protobuf CreateVehicleRequest
	PBreq := &pb.RentVehicleResponse{
		Id:              "",
		Unicode:         vehicle.Unicode,
		RequiredExperts: vehicle.RequiredExperts,
		Speed:           vehicle.Speed,
		RentPrice:       vehicle.RentPrice,
		IsActive:        vehicle.IsActive,
		Type:            vehicle.Type,
		Owner:           &pb.Owner{Id: vehicle.Owner.Id, FirstName: vehicle.Owner.FirstName, LastName: vehicle.Owner.LastName, Email: vehicle.Owner.Email},
		Passenger:       int32(vehicle.Passenger),
		Model:           int32(vehicle.Model),
	}

	return PBreq, nil
}

