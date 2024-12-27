package vehicle

import (
	"errors"

	"github.com/onlineTraveling/vehicle/api/pb"
	"github.com/onlineTraveling/vehicle/internal/vehicle/domain"
)

func PBVehicleRequest2DomainVehicle(request interface{}) (domain.Vehicle, error) {
    var PBreq struct {
        Id              string
        Unicode         string
        RequiredExperts int32
        Speed           int32
        RentPrice       int32
        IsActive        bool
        Type            string
        Owner           *struct {
            Id        uint64
            FirstName string
            LastName  string
            Email     string
        }
    }

    // Map the input to the common structure
    switch req := request.(type) {
    case *pb.CreateVehicleRequest:
        PBreq = struct {
            Id              string
            Unicode         string
            RequiredExperts int32
            Speed           int32
            RentPrice       int32
            IsActive        bool
            Type            string
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
    case *pb.UpdateVehicleRequest:
        PBreq = struct {
            Id              string
            Unicode         string
            RequiredExperts int32
            Speed           int32
            RentPrice       int32
            IsActive        bool
            Type            string
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
        return domain.Vehicle{}, errors.New("unsupported request type")
    }

    // Validate input
    if PBreq.Owner == nil {
        return domain.Vehicle{}, errors.New("OWNER CANNOT BE NIL")
    }
    if PBreq.Unicode == "" {
        return domain.Vehicle{}, errors.New("UNICODE CANNOT BE EMPTY")
    }
    if PBreq.Owner.Id == 0 || PBreq.Owner.FirstName == "" || PBreq.Owner.LastName == "" || PBreq.Owner.Email == "" {
        return domain.Vehicle{}, errors.New("OWNER DETAILS ARE INCOMPLETE")
    }

    // Construct and return the domain vehicle
    vehicle := domain.Vehicle{
        Id:              PBreq.Id,
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
		Owner: &pb.Owner{
			Id:        vehicle.Owner.Id,
			FirstName: vehicle.Owner.FirstName,
			LastName:  vehicle.Owner.LastName,
			Email:     vehicle.Owner.Email,
		},
	}

	return PBreq, nil
}
