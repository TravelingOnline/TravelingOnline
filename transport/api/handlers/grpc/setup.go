package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/onlineTraveling/transport/api/handlers/grpc/tour"
	"github.com/onlineTraveling/transport/api/handlers/grpc/company"
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
	c:= company.NewGRPCTransportHandler(*cHandler)
	t:= tour.NewGRPCTourHandler(*tHandler)
	// pb.RegisterTrasportServiceServer(grpcServer, d)
	pb.RegisterCompanyServiceServer(grpcServer, c)
	pb.RegisterTourServiceServer(grpcServer, t)

	log.Printf("Server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
