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
	conn, err := grpc.Dial("0.0.0.0:50052", grpc.WithInsecure())
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
		SenderOwnerID:   "e5824ec0-48ae-4712-9178-48f96e33328c",
		ReceiverOwnerID: "f2fc63c0-443e-4181-b0c9-2316a4a9845c",
		Amount:          120,
	}

	// Call the CreateWallet method
	response, err := client.Transfer(ctx, in)
	if err != nil {
		log.Fatalf("cannot transfer: %v", err)
	}

	// Map the response to the domain model
	_, err = mappers.TransferEntitieToTransferDomain(response)
	if err != nil {
		log.Fatalf("cannot map response: %v", err)
	}

	// Print the domain response
	log.Printf("Transfered successfully")
}
