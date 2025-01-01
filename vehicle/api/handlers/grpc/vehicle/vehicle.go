package vehicle

import (
	"context"
	"fmt"

	"github.com/onlineTraveling/vehicle/api/pb"
	"github.com/onlineTraveling/vehicle/api/service"
	"github.com/onlineTraveling/vehicle/internal/vehicle/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCVehicleHandler struct {
	pb.UnimplementedVehicleServiceServer
	vehicleService *service.VehicleService
}

func NewGRPCVehicleHandler(vehicleService service.VehicleService) *GRPCVehicleHandler {
	return &GRPCVehicleHandler{vehicleService: &vehicleService}
}

func (g *GRPCVehicleHandler) CreateVehicle(ctx context.Context, req *pb.CreateVehicleRequest) (*pb.CreateVehicleResponse, error) {
	domainRequest, err := PBVehicleRequest2DomainVehicle(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	vID, err := g.vehicleService.CreateVehicle(ctx, &domainRequest)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.CreateVehicleResponse{
		Id: string(*vID),
	}, nil
}

func (g *GRPCVehicleHandler) UpdateVehicle(ctx context.Context, req *pb.UpdateVehicleRequest) (*pb.UpdateVehicleResponse, error) {
	domainRequest, err := PBVehicleRequest2DomainVehicle(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	vID, err := g.vehicleService.UpdateVehicle(ctx, &domainRequest)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.UpdateVehicleResponse{
		Id: string(*vID),
	}, nil
}

func (g *GRPCVehicleHandler) DeleteVehicle(ctx context.Context, delReq *pb.DeleteVehicleRequest) (*pb.DeleteVehicleResponse, error) {
	vID := domain.VehicleID(delReq.Id)
	deletedVehcileID, err := g.vehicleService.DeleteVehicle(ctx, &vID)
	if err != nil {
		return deletedVehcileID, status.Errorf(codes.Internal, err.Error())
	}
	return deletedVehcileID, nil
}

func (g *GRPCVehicleHandler) GetVehicle(ctx context.Context, vehicleReq *pb.GetVehicleRequest) (*pb.GetVehicleResponse, error) {
	vID := domain.VehicleID(vehicleReq.Id)
	vehicle, err := g.vehicleService.GetVehicle(ctx, &vID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	domainVehicle, err := DomainVehicle2PBVehicleResponse(*vehicle)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	return domainVehicle, nil
}

func (g *GRPCVehicleHandler) RentVehicle(ctx context.Context, req *pb.RentVehicleRequest) (*pb.RentVehicleResponse, error) {

	rentReq, err := PBVehicleRequest2DomainVehicle(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	vehicle, err := g.vehicleService.RentVehicle(ctx, rentReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	bestVehicle, err := DomainVehicle2PBRentVehicleRequest(*vehicle)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	return bestVehicle, nil
}
