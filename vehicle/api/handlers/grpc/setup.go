package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/onlineTraveling/vehicle/api/pb"
	"github.com/onlineTraveling/vehicle/app"
	"github.com/onlineTraveling/vehicle/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedVehicleServiceServer
}

func NewGRPCServer() pb.VehicleServiceServer {
	return new(Server)
}

func Run(cfg config.Config, app *app.App) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Server.HttpPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	// auth.RegisterBankServiceServer(s, handlers.NewGRPCBankHandler(app.BankService()))
	// Register the Health Service server
	// healthServer := &vehicle.GRPCServer{}
	pb.RegisterVehicleServiceServer(s, NewGRPCServer())

	// Register reflection service on gRPC server
	reflection.Register(s)

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
