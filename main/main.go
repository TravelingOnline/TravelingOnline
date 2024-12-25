package main

import (
	"log"
	"net"

	"github.com/onlineTraveling/bank/protobufs" // Update to your actual import path
	"google.golang.org/grpc"
)

type bankServiceServer struct {
	protobufs.UnimplementedBankServiceServer
}

// Implement the methods for BankServiceServer interface (e.g., CreateWallet, Transfer)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	protobufs.RegisterBankServiceServer(grpcServer, &bankServiceServer{}) // Correct service registration

	log.Println("Bank service running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
