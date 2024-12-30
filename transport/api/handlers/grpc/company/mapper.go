package company

import (
	"errors"

	"github.com/onlineTraveling/transport/api/pb"
	"github.com/onlineTraveling/transport/internal/company/domain"
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

	// Construct and return the domain company
	company := domain.Company{
		Id:    PBreq.Id,
		Name:  PBreq.Name,
		Owner: &domain.Owner{Id: PBreq.Owner.Id, FirstName: PBreq.Owner.FirstName, LastName: PBreq.Owner.LastName, Email: PBreq.Owner.Email},
	}

	return company, nil
}

func DomainCompany2PBCreateCompanyRequest(company domain.Company) (*pb.CreateCompanyRequest, error) {
	// Validate input
	if company.Owner == nil {
		return nil, errors.New("company.Owner cannot be nil")
	}

	// Construct and return the protobuf CreateCompanyRequest
	PBreq := &pb.CreateCompanyRequest{
		Id:    "",
		Name:  company.Name,
		Owner: &pb.Owner{Id: company.Owner.Id, FirstName: company.Owner.FirstName, LastName: company.Owner.LastName, Email: company.Owner.Email},
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
		Id:   company.Id,
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
