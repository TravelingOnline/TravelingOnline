package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/onlineTraveling/vehicle/api/handlers/grpc/vehicle"
	"github.com/onlineTraveling/vehicle/api/pb"
	"github.com/onlineTraveling/vehicle/app"
	"github.com/onlineTraveling/vehicle/config"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedVehicleServiceServer
}

func Run(cfg config.Config, app *app.App) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Server.HttpPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterVehicleServiceServer(grpcServer, vehicle.NewVehicleGRPCServer())

	log.Println("Server is running on port 8081...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
