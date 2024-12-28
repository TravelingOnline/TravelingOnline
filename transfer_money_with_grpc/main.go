package main

import (
	"context"
	"log"

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
	in := &protobufs.TransferRequest{
		SenderOwnerID:   "bd8873c8-be31-4285-b8ef-b7ed2ac82be8",
		ReceiverOwnerID: "5f664e09-e970-48c2-8fc2-dceb3096bd52",
		Amount:          120,
	}

	// Call the CreateWallet method
	response, err := client.Transfer(ctx, in)
	if err != nil {
		log.Fatalf("cannot transfer: %v", err)
	}

	// Map the response to the domain model
	domainResponse, err := mappers.TransferEntitieToTransferDomain(response)
	if err != nil {
		log.Fatalf("cannot map response: %v", err)
	}

	// Print the domain response
	log.Printf("transfered successfully: %v", *domainResponse)
}
