package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/onlineTraveling/bank/api/grpc/handlers"
	"github.com/onlineTraveling/bank/app"
	"github.com/onlineTraveling/bank/config"

	"github.com/onlineTraveling/bank/protobufs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func Run(cfg config.Config, app *app.App) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Server.GRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	healthServer := &handlers.HealthServer{}
	grpc_health_v1.RegisterHealthServer(s, healthServer)

	reflection.Register(s)

	log.Println("GRPC server started..")

	bankHandler := handlers.NewGRPCBankHandler(app.BankService())
	protobufs.RegisterBankServiceServer(s, bankHandler)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
