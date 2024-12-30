package tour


import (
	"errors"

	"github.com/onlineTraveling/transport/api/pb"
	"github.com/onlineTraveling/transport/internal/tour/domain"
)

func PBTourRequest2DomainTour(request interface{}) (domain.Tour, error) {
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
	case *pb.CreateTourRequest:
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
	case *pb.UpdateTourRequest:
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
		return domain.Tour{}, errors.New("unsupported request type")
	}

	// Validate input
	if PBreq.Owner == nil {
		return domain.Tour{}, errors.New("OWNER CANNOT BE NIL")
	}
	if PBreq.Owner.Id == 0 || PBreq.Owner.FirstName == "" || PBreq.Owner.LastName == "" || PBreq.Owner.Email == "" {
		return domain.Tour{}, errors.New("OWNER DETAILS ARE INCOMPLETE")
	}

	// Construct and return the domain company
	company := domain.Tour{
		Id:    PBreq.Id,
		Name:  PBreq.Name,
		Owner: &domain.Owner{Id: PBreq.Owner.Id, FirstName: PBreq.Owner.FirstName, LastName: PBreq.Owner.LastName, Email: PBreq.Owner.Email},
	}

	return company, nil
}

func DomainCompany2PBCreateCompanyRequest(tour domain.Tour) (*pb.CreateCompanyRequest, error) {
	// Validate input
	if tour.Owner == nil {
		return nil, errors.New("company.Owner cannot be nil")
	}

	// Construct and return the protobuf CreateCompanyRequest
	PBreq := &pb.CreateCompanyRequest{
		Id:    "",
		Name:  tour.Name,
		Owner: &pb.Owner{Id: tour.Owner.Id, FirstName: tour.Owner.FirstName, LastName: tour.Owner.LastName, Email: tour.Owner.Email},
	}

	return PBreq, nil
}

func DomainCompany2PBCompanyResponse(tour domain.Tour) (*pb.GetByIDTourResponse, error) {
	// Validate input
	if tour.Owner == nil {
		return nil, errors.New("company.Owner cannot be nil")
	}

	// Construct and return the protobuf CreateCompanyResponse
	PBreq := &pb.GetByIDTourResponse{
		Id:   tour.Id,
		Name: tour.Name,
		Owner: &pb.Owner{
			Id:        tour.Owner.Id,
			FirstName: tour.Owner.FirstName,
			LastName:  tour.Owner.LastName,
			Email:     tour.Owner.Email,
		},
	}

	return PBreq, nil
}
