package vehicle

import (
	"context"

	"github.com/onlineTraveling/vehicle/api/pb"
	"github.com/onlineTraveling/vehicle/api/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

type GRPCVehicleHandler struct {
	pb.UnimplementedVehicleServiceServer
	vehicleService *service.VehicleService
}

func NewGRPCVehicleHandler(vehicleService *service.VehicleService) *GRPCVehicleHandler {
	return &GRPCVehicleHandler{vehicleService: vehicleService}
}

func (g *GRPCVehicleHandler) CreateVehicle(ctx context.Context, req *pb.CreateVehicleRequest) (*pb.CreateVehicleResponse, error) {
	domainRequest, err := PBCreateVehicleRequest2DomainVehicle(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	_, err = g.vehicleService.CreateVehicle(ctx, &domainRequest)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.CreateVehicleResponse{
		Id: "1111111",
	}, nil
}

type GRPCServer struct {
	pb.UnimplementedVehicleServiceServer
}

// Check implements Health.Check
func (s *GRPCServer) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}
