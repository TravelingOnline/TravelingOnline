package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/onlineTraveling/transport/api/pb"
	"github.com/onlineTraveling/transport/internal/tour/domain"
	"github.com/onlineTraveling/transport/internal/tour/port"
)

type TourService struct {
	srv port.Service
}

func NewTourService(srv port.Service) *TourService {
	return &TourService{
		srv: srv,
	}
}

func (v *TourService) CreateTour(ctx context.Context, req *domain.Tour) (*domain.TourID, error) {
	req.Id = uuid.New().String()
	tID, err := v.srv.CreateTourService(ctx, *req)
	if err != nil {
		return nil, err
	}
	return &tID, nil
}

func (v *TourService) UpdateTour(ctx context.Context, req *domain.Tour) (*domain.TourID, error) {
	tID, err := v.srv.UpdateTourService(ctx, *req)
	if err != nil {
		return nil, err
	}
	return &tID, nil
}

func (v *TourService) DeleteTour(ctx context.Context, vID *domain.TourID) (*pb.DeleteTourResponse, error) {
	// Call the service to delete the tour
	deletedTourID, err := v.srv.DeleteTourService(ctx, *vID)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteTourResponse{
		Id: string(deletedTourID),
	}, nil
}

func (v *TourService) GetByIDTour(ctx context.Context, vID *domain.TourID) (*domain.Tour, error) {
	tour, err := v.srv.GetByIDTourService(ctx, *vID)
	if err != nil {
		return &domain.Tour{}, err
	}
	return &tour, nil
}
