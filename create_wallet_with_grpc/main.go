package main

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/onlineTraveling/bank/protobufs"
	"google.golang.org/grpc"
)

func main() {
	// Define the address of the bank service
	bankServiceAddress := "localhost:50051" // Replace with the actual address of the bank service

	// Establish a connection to the bank service
	conn, err := grpc.Dial(bankServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to bank service: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client for the bank service
	bankClient := protobufs.NewBankServiceClient(conn)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a wallet request
	createWalletReq := &protobufs.CreateWalletRequest{
		UserID: uuid.New().String(), // Replace with a valid UUID
	}

	// Call the CreateWallet RPC
	resp, err := bankClient.CreateWallet(ctx, createWalletReq)
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	// Print the response
	log.Printf("Wallet created successfully: %s", resp.Message)
}
