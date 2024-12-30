package tour

import (
	"context"

	"github.com/onlineTraveling/transport/api/pb"
	"github.com/onlineTraveling/transport/api/service"
	"github.com/onlineTraveling/transport/internal/tour/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCTourHandler struct {
	pb.UnimplementedTourServiceServer
	tourService *service.TourService
}

func NewGRPCTourHandler(tourService service.TourService) *GRPCTourHandler {
	return &GRPCTourHandler{tourService: &tourService}
}

func (g *GRPCTourHandler) CreateTour(ctx context.Context, req *pb.CreateTourRequest) (*pb.CreateTourResponse, error) {
	domainRequest, err := PBTourRequest2DomainTour(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	tID, err := g.tourService.CreateTour(ctx, &domainRequest)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.CreateTourResponse{
		Id: string(*tID),
	}, nil
}

func (g *GRPCTourHandler) UpdateTour(ctx context.Context, req *pb.UpdateTourRequest) (*pb.UpdateTourResponse, error) {
	domainRequest, err := PBTourRequest2DomainTour(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	vID, err := g.tourService.UpdateTour(ctx, &domainRequest)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.UpdateTourResponse{
		Id: string(*vID),
	}, nil
}

func (g *GRPCTourHandler) DeleteTour(ctx context.Context, delReq *pb.DeleteTourRequest) (*pb.DeleteTourResponse, error) {
	vID := domain.TourID(delReq.Id)
	deletedTourID, err := g.tourService.DeleteTour(ctx, &vID)
	if err != nil {
		return deletedTourID, status.Errorf(codes.Internal, err.Error())
	}
	return deletedTourID, nil
}

func (g *GRPCTourHandler) GetByIDTour(ctx context.Context, TourReq *pb.GetByIDTourRequest) (*pb.GetByIDTourResponse, error) {
	tID := domain.TourID(TourReq.Id)
	tour, err := g.tourService.GetByIDTour(ctx, &tID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	domaintour, err := DomainTour2PBTourResponse(*tour)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	return domaintour, nil
}
