package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/onlineTraveling/transport/api/handlers/grpc/transport"
	"github.com/onlineTraveling/transport/api/pb"
	"github.com/onlineTraveling/transport/api/service"
	"github.com/onlineTraveling/transport/app"
	"github.com/onlineTraveling/transport/config"
	"google.golang.org/grpc"
)

func Run(cfg config.Config, app *app.App) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Server.HttpPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	cHandler := service.NewCompanyService(app.CompanyService())
	tHandler := service.NewTourService(app.TourService())
	d := transport.NewGRPCTransportHandler(*cHandler)
	// pb.RegisterTrasportServiceServer(grpcServer, d)
	pb.RegisterCompanyServiceServer(grpcServer, d)
	pb.RegisterTourServiceServer(grpcServer, d)

	log.Printf("Server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
