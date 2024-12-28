package transport

import (
	"errors"

	"github.com/onlineTraveling/transport/api/pb"
	"github.com/onlineTraveling/transport/internal/transport/domain"
)

func PBCompanyRequest2DomainCompany(request interface{}) (domain.Company, error) {
	var PBreq struct {
		Id    string
		Name  string
		Owner *struct {
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
			Id    string
			Name  string
			Owner *struct {
				Id        uint64
				FirstName string
				LastName  string
				Email     string
			}
		}{
			Id:   req.Id,
			Name: req.Name,
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
			Id    string
			Name  string
			Owner *struct {
				Id        uint64
				FirstName string
				LastName  string
				Email     string
			}
		}{
			Id:   req.Id,
			Name: req.Name,

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
	if PBreq.Owner.Id == 0 || PBreq.Owner.FirstName == "" || PBreq.Owner.LastName == "" || PBreq.Owner.Email == "" {
		return domain.Company{}, errors.New("OWNER DETAILS ARE INCOMPLETE")
	}

	// Construct and return the domain vehicle
	vehicle := domain.Company{
		Id:    PBreq.Id,
		Name:  PBreq.Name,
		Owner: &domain.Owner{Id: PBreq.Owner.Id, FirstName: PBreq.Owner.FirstName, LastName: PBreq.Owner.LastName, Email: PBreq.Owner.Email},
	}

	return vehicle, nil
}

func DomainVehicle2PBCreateVehicleRequest(vehicle domain.Company) (*pb.CreateCompanyRequest, error) {
	// Validate input
	if vehicle.Owner == nil {
		return nil, errors.New("vehicle.Owner cannot be nil")
	}

	// Construct and return the protobuf CreateVehicleRequest
	PBreq := &pb.CreateCompanyRequest{
		Id:    "",
		Name:  vehicle.Name,
		Owner: &pb.Owner{Id: vehicle.Owner.Id, FirstName: vehicle.Owner.FirstName, LastName: vehicle.Owner.LastName, Email: vehicle.Owner.Email},
	}

	return PBreq, nil
}

func DomainCompany2PBCompanyResponse(company domain.Company) (*pb.GetByIDCompanyResponse, error) {
	// Validate input
	if company.Owner == nil {
		return nil, errors.New("company.Owner cannot be nil")
	}

	// Construct and return the protobuf CreateCompanyResponse
	PBreq := &pb.GetByIDCompanyResponse{
		Name: company.Name,
		Owner: &pb.Owner{
			Id:        company.Owner.Id,
			FirstName: company.Owner.FirstName,
			LastName:  company.Owner.LastName,
			Email:     company.Owner.Email,
		},
	}

	return PBreq, nil
}
