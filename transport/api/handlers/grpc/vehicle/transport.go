package transport

import (
	"context"

	"github.com/onlineTraveling/transport/api/pb"
	"github.com/onlineTraveling/transport/api/service"
	"github.com/onlineTraveling/transport/internal/transport/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCTransportHandler struct {
	pb.UnimplementedTrasportServiceServer
	trasnportService *service.TransportService
}

func NewGRPCTransportHandler(trasnsportService service.TransportService) *GRPCTransportHandler{
	return &GRPCTransportHandler{trasnportService: &trasnsportService}
}

func (g *GRPCTransportHandler) CreateVehicle(ctx context.Context, req *pb.CreateCompanyRequest) (*pb.CreateCompanyResponse, error) {
	domainRequest, err := PBVehicleRequest2DomainVehicle(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	vID, err := g.vehicleService.CreateVehicle(ctx, &domainRequest)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	// log.Println(">>>>>>>>>>>>", string(*vID))
	return &pb.CreateVehicleResponse{
		Id: string(*vID),
	}, nil
}

func (g *GRPCTransportHandler) UpdateVehicle(ctx context.Context, req *pb.UpdateCompanyRequest) (*pb.UpdateCompanyResponse, error) {
	domainRequest, err := PBVehicleRequest2DomainVehicle(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	vID, err := g.vehicleService.UpdateVehicle(ctx, &domainRequest)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	// log.Println(">>>>>>>>>>>>", string(*vID))
	return &pb.UpdateVehicleResponse{
		Id: string(*vID),
	}, nil
}

func (g *GRPCTransportHandler) DeleteVehicle(ctx context.Context, delReq *pb.DeleteCompanyRequest) (*pb.DeleteCompanyResponse, error) {
	vID := domain.VehicleID(delReq.Id)
	deletedVehcileID, err := g.vehicleService.DeleteVehicle(ctx, &vID)
	if err != nil {
		return deletedVehcileID, status.Errorf(codes.Internal, err.Error())
	}
	return deletedVehcileID, nil
}

func (g *GRPCTransportHandler) GetVehicle(ctx context.Context, vehicleReq *pb.GetByIDCompanyRequest) (*pb.GetByIDCompanyResponse, error) {
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
