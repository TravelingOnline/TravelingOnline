package tour

import (
	"errors"
	"log"

	"github.com/onlineTraveling/transport/api/pb"
	"github.com/onlineTraveling/transport/internal/tour/domain"
)

func PBTourRequest2DomainTour(request interface{}) (domain.Tour, error) {
	var domainTour domain.Tour

	switch req := request.(type) {
	case *pb.CreateTourRequest:
		domainTour = mapCreateTourRequest(req)
	case *pb.UpdateTourRequest:
		domainTour = mapUpdateTourRequest(req)
	default:
		log.Printf("unsupported request type: %T", request)
		return domain.Tour{}, errors.New("unsupported request type")
	}

	return domainTour, nil
}

func mapUpdateTourRequest(req *pb.UpdateTourRequest) domain.Tour {
	return domain.Tour{
		Id:           req.Id,
		Source:       req.Source,
		Destination:  req.Destination,
		StartDate:    req.StartDate,
		EndDate:      req.EndDate,
		Type:         req.Type,
		Price:        req.Price,
		Capacity:     req.Capacity,
		AdminApprove: req.AdminApprove,
		Ended:        req.Ended,
		CompanyID:    req.CompanyID,
		Vehicle: func() domain.Vehicle {
			if req.Vehicle != nil {
				return domain.Vehicle{
					Id:              req.Vehicle.Id,
					Unicode:         req.Vehicle.Unicode,
					RequiredExperts: req.Vehicle.RequiredExperts,
					Speed:           req.Vehicle.Speed,
					RentPrice:       req.Vehicle.RentPrice,
					Type:            req.Vehicle.Type,
					Passenger:       req.Vehicle.Passenger,
					Model:           req.Vehicle.Model,
				}
			}
			return domain.Vehicle{} // Return an empty Vehicle if none is provided
		}(),
		TechnicalTeam: mapTechnicalTeam(req.TechnicalTeam),
	}
}

func mapCreateTourRequest(req *pb.CreateTourRequest) domain.Tour {
	return domain.Tour{
		Id:           req.Id,
		Source:       req.Source,
		Destination:  req.Destination,
		StartDate:    req.StartDate,
		EndDate:      req.EndDate,
		Type:         req.Type,
		Price:        req.Price,
		Capacity:     req.Capacity,
		AdminApprove: false,
		Ended:        false,
		CompanyID:    req.CompanyID,
		Vehicle: func() domain.Vehicle {
			if req.Vehicle != nil {
				return domain.Vehicle{
					Id:              req.Vehicle.Id,
					Unicode:         req.Vehicle.Unicode,
					RequiredExperts: req.Vehicle.RequiredExperts,
					Speed:           req.Vehicle.Speed,
					RentPrice:       req.Vehicle.RentPrice,
					Type:            req.Vehicle.Type,
					Passenger:       req.Vehicle.Passenger,
					Model:           req.Vehicle.Model,
				}
			}
			return domain.Vehicle{}
		}(),
		TechnicalTeam: mapTechnicalTeam(req.TechnicalTeam),
	}
}

func mapTechnicalTeam(teams []*pb.TechnicalTeam) []*domain.TechnicalTeam {
	result := make([]*domain.TechnicalTeam, len(teams))
	for i, t := range teams {
		result[i] = &domain.TechnicalTeam{
			Id:        t.Id,
			FirstName: t.FirstName,
			LastName:  t.LastName,
			Age:       t.Age,
			Expertise: t.Expertise,
		}
	}

	return result
}

func DomainTour2PBCreateTourRequest(tour domain.Tour) (*pb.CreateTourRequest, error) {
	// Map the domain.Tour to pb.CreateTourRequest

	PBreq := &pb.CreateTourRequest{
		Id:          tour.Id,
		Source:      tour.Source,
		Destination: tour.Destination,
		StartDate:   tour.StartDate,
		EndDate:     tour.EndDate,
		Type:        tour.Type,
		Price:       tour.Price,
		Capacity:    tour.Capacity,
		CompanyID:   tour.CompanyID,
		Vehicle: &pb.Vehicle{
			Id:              tour.Vehicle.Id,
			Unicode:         tour.Vehicle.Unicode,
			RequiredExperts: tour.Vehicle.RequiredExperts,
			Speed:           tour.Vehicle.Speed,
			RentPrice:       tour.Vehicle.RentPrice,
			Type:            tour.Vehicle.Type,
			Passenger:       tour.Vehicle.Passenger,
			Model:           tour.Vehicle.Model,
		},
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
		Id:           tour.Id,
		Source:       tour.Source,
		Destination:  tour.Destination,
		StartDate:    tour.StartDate,
		EndDate:      tour.EndDate,
		Type:         tour.Type,
		Price:        tour.Price,
		Capacity:     tour.Capacity,
		Ended:        tour.Ended,
		AdminApprove: tour.AdminApprove,
		CompanyID:    tour.CompanyID,
		Vehicle: &pb.Vehicle{
			Id:              tour.Vehicle.Id,
			Unicode:         tour.Vehicle.Unicode,
			RequiredExperts: tour.Vehicle.RequiredExperts,
			Speed:           tour.Vehicle.Speed,
			RentPrice:       tour.Vehicle.RentPrice,
			Type:            tour.Vehicle.Type,
			Passenger:       tour.Vehicle.Passenger,
			Model:           tour.Vehicle.Model,
		},
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
