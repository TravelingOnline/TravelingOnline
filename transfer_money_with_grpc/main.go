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
		SenderOwnerID:   "f3e6b4d6-7887-4214-b3e7-04afb7e1e6be",
		ReceiverOwnerID: "4e09cacc-eb58-4796-a9d8-e678034a484c",
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
