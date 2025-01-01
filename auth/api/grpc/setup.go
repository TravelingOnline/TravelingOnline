package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/onlineTraveling/auth/api/grpc/handlers"
	"github.com/onlineTraveling/auth/app"
	"github.com/onlineTraveling/auth/config"

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

	log.Println("Auth GRPC server started..")

	// userHandler := handlers.NewGRPCUserHandler(app.UserService(context.Background()))
	// protobufs.RegisterAuthServiceServer(s, userHandler)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
