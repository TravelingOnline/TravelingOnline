package service

import (
	"agency/api/pb"
	"agency/internal/tour"
	"agency/internal/tour/domain"
	tourPort "agency/internal/tour/port"
	"context"
)

type TourService struct {
	svc tourPort.Service
}

func NewTourService(svc tourPort.Service) *TourService {
	return &TourService{svc}
}

var (
	ErrTourOnCreate = tour.ErrTourOnCreate
	ErrTourOnUpdate = tour.ErrTourOnUpdate
	ErrTourNotFound = tour.ErrTourNotFound
)

func (s *TourService) CreateTour(ctx context.Context, req *pb.TourCreateRequest) error {
	err := s.svc.CreateTour(ctx, domain.Tour{
		Capacity:         uint(req.GetCapacity()),
		Price:            req.GetPrice(),
		IsActive:         req.GetIsActive(),
		OutboundTicketID: uint(req.GetOutboundTicketId()),
		ReturnTicketID:   uint(req.GetReturnTicketId()),
		HotelID:          uint(req.GetHotelId()),
	})

	if err != nil {
		return ErrTourOnCreate
	}

	return nil
}
