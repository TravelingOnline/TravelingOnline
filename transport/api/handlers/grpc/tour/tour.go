package tour

import (
	"context"

	"github.com/onlineTraveling/transport/api/pb"
	"github.com/onlineTraveling/transport/api/service"
	"github.com/onlineTraveling/transport/internal/tour/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCTransportHandler struct {
	pb.UnimplementedTourServiceServer
	tourService *service.TourService
}

func NewGRPCTransportHandler(tourService service.TourService) *GRPCTransportHandler {
	return &GRPCTransportHandler{tourService: &tourService}
}

func (g *GRPCTransportHandler) CreateTour(ctx context.Context, req *pb.CreateTourRequest) (*pb.CreateTourResponse, error) {
	domainRequest, err := PBCompanyRequest2DomainCompany(req)
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

func (g *GRPCTransportHandler) UpdateTour(ctx context.Context, req *pb.UpdateCompanyRequest) (*pb.UpdateTourResponse, error) {
	domainRequest, err := PBCompanyRequest2DomainCompany(req)
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

func (g *GRPCTransportHandler) DeleteTour(ctx context.Context, delReq *pb.DeleteCompanyRequest) (*pb.DeleteTourResponse, error) {
	vID := domain.TourID(delReq.Id)
	deletedTourID, err := g.tourService.DeleteTour(ctx, &vID)
	if err != nil {
		return deletedTourID, status.Errorf(codes.Internal, err.Error())
	}
	return deletedTourID, nil
}

func (g *GRPCTransportHandler) GetByIDTour(ctx context.Context, CompanyReq *pb.GetByIDTourRequest) (*pb.GetByIDTourResponse, error) {
	tID := domain.TourID(CompanyReq.Id)
	tour, err := g.tourService.GetByIDTour(ctx, &tID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	domaintour, err := DomainCompany2PBCompanyResponse(*tour)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	return domaintour, nil
}
