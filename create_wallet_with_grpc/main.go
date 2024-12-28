package main

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/onlineTraveling/bank/pkg/adapters/clients/grpc/mappers"
	"github.com/onlineTraveling/bank/protobufs"
	"google.golang.org/grpc"
)

func main() {
	// Create a gRPC connection to the bank service
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("client cannot connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a new BankService client
	client := protobufs.NewBankServiceClient(conn)

	// Create a context
	ctx := context.Background()

	// Prepare the request
	in := &protobufs.CreateWalletRequest{
		UserID: uuid.New().String(),
	}

	// Call the CreateWallet method
	response, err := client.CreateWallet(ctx, in)
	if err != nil {
		log.Fatalf("cannot create wallet: %v", err)
	}

	// Map the response to the domain model
	domainResponse, err := mappers.CreateWalletResponseToMessageDomain(response)
	if err != nil {
		log.Fatalf("cannot map response: %v", err)
	}

	// Print the domain response
	log.Printf("Wallet created successfully: %v", *domainResponse)
}
