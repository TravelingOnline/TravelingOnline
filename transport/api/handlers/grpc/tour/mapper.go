package tour

import (
	"errors"

	"github.com/onlineTraveling/transport/api/pb"
	"github.com/onlineTraveling/transport/internal/tour/domain"
)

func PBTourRequest2DomainTour(request interface{}) (domain.Tour, error) {
	var domainTour domain.Tour

	switch req := request.(type) {
	case *pb.CreateTourRequest:
		// Map the CreateTourRequest fields to the domain.Tour struct
		domainTour = domain.Tour{
			Id:             req.Id,
			Source:         req.Source,
			Destination:    req.Destination,
			StartDate:      req.StartDate,
			EndDate:        req.EndDate,
			Type:           req.Type,
			Price:          req.Price,
			VehicleUnicode: req.VehicleUnicode,
			TechnicalTeam: func() []*domain.TechnicalTeam {
				team := make([]*domain.TechnicalTeam, len(req.TechnicalTeam))
				for i, t := range req.TechnicalTeam {
					team[i] = &domain.TechnicalTeam{
						Id:        t.Id,
						FirstName: t.FirstName,
						LastName:  t.LastName,
						Age:       t.Age,
						Expertise: t.Expertise,
					}
				}
				return team
			}(),
		}
	default:
		return domain.Tour{}, errors.New("unsupported request type")
	}

	return domainTour, nil
}

func DomainTour2PBCreateTourRequest(tour domain.Tour) (*pb.CreateTourRequest, error) {
	// Map the domain.Tour to pb.CreateTourRequest
	PBreq := &pb.CreateTourRequest{
		Id:             tour.Id,
		Source:         tour.Source,
		Destination:    tour.Destination,
		StartDate:      tour.StartDate,
		EndDate:        tour.EndDate,
		Type:           tour.Type,
		Price:          tour.Price,
		VehicleUnicode: tour.VehicleUnicode,
		TechnicalTeam: func() []*pb.TechnicalTeam {
			team := make([]*pb.TechnicalTeam, len(tour.TechnicalTeam))
			for i, t := range tour.TechnicalTeam {
				team[i] = &pb.TechnicalTeam{
					Id:        t.Id,
					FirstName: t.FirstName,
					LastName:  t.LastName,
					Age:       t.Age,
					Expertise: t.Expertise,
				}
			}
			return team
		}(),
	}

	return PBreq, nil
}


func DomainTour2PBTourResponse(tour domain.Tour) (*pb.GetByIDTourResponse, error) {
	// Map the domain.Tour to pb.GetByIDTourResponse
	PBresp := &pb.GetByIDTourResponse{
		Id:             tour.Id,
		Source:         tour.Source,
		Destination:    tour.Destination,
		StartDate:      tour.StartDate,
		EndDate:        tour.EndDate,
		Type:           tour.Type,
		Price:          tour.Price,
		VehicleUnicode: tour.VehicleUnicode,
		TechnicalTeam: func() []*pb.TechnicalTeam {
			team := make([]*pb.TechnicalTeam, len(tour.TechnicalTeam))
			for i, t := range tour.TechnicalTeam {
				team[i] = &pb.TechnicalTeam{
					Id:        t.Id,
					FirstName: t.FirstName,
					LastName:  t.LastName,
					Age:       t.Age,
					Expertise: t.Expertise,
				}
			}
			return team
		}(),
	}

	return PBresp, nil
}

