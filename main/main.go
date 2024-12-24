package main

import (
	"context"
	"log"
	"time"

	"github.com/onlineTraveling/bank/protobufs"
	"google.golang.org/grpc"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure()) // Replace with actual gRPC server address
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a client for the BankService
	bankService := protobufs.NewBankServiceClient(conn)

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Call the CreateWallet method
	req := &protobufs.CreateWalletRequest{
		// Fill in the request fields as per your .proto file
		UserID: "12345",
	}

	resp, err := bankService.CreateWallet(ctx, req)
	if err != nil {
		log.Fatalf("Error calling CreateWallet: %v", err)
	}

	log.Printf("CreateWallet Response: %v", resp)
}
