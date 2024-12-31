package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/onlineTraveling/vehicle/api/handlers/grpc/vehicle"
	"github.com/onlineTraveling/vehicle/api/pb"
	"github.com/onlineTraveling/vehicle/api/service"
	"github.com/onlineTraveling/vehicle/app"
	"github.com/onlineTraveling/vehicle/config"
	"google.golang.org/grpc"
)

func Run(cfg config.Config, app *app.App) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Server.HttpPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	vHandler := service.NewVehicleService(app.VehicleService())
	d := vehicle.NewGRPCVehicleHandler(*vHandler)
	pb.RegisterVehicleServiceServer(grpcServer, d)


	log.Printf("Server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
