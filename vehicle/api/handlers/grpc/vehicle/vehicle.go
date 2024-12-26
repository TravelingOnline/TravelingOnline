package vehicle

import (
	"context"
	"log"

	"github.com/onlineTraveling/vehicle/api/pb"
	"github.com/onlineTraveling/vehicle/api/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type vehicleGRPCApi struct {
	createVehicleService *service.VehicleService
	pb.UnimplementedVehicleServiceServer
}

func NewVehicleGRPCServer() pb.VehicleServiceServer {
	return new(vehicleGRPCApi)
}

func (s *vehicleGRPCApi) CreateVehicle(ctx context.Context, v *pb.CreateVehicleRequest) (*pb.CreateVehicleResponse, error) {
	// Map the protobuf request to the domain model
	newVehicle, err := PBCreateVehicleRequest2DomainVehicle(v)
	if err != nil {
		log.Printf("error mapping CreateVehicleRequest to domain model: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	// Call the service to create the vehicle
	vehicleID, err := s.createVehicleService.CreateVehicle(ctx, &newVehicle)
	if err != nil {
		log.Printf("error creating vehicle: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to create vehicle: %v", err)
	}

	// Successfully created vehicle, return response
	return &vehicleID, nil
}
