package helpers

import (
	"fmt"
	"time"

	vpb "github.com/onlineTraveling/transport/pkg/adapters/storage/vehicle-pb"
	bpb "github.com/onlineTraveling/transport/pkg/adapters/storage/bank-pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewVehicleClient(host *string, port *uint) (vpb.VehicleServiceClient, *grpc.ClientConn, error) {
	// Dial the gRPC server using the provided configuration
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", *host, *port), // Format HttpPort as an int
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to VehicleService: %w", err)
	}

	client := vpb.NewVehicleServiceClient(conn)
	return client, conn, nil
}

func NewBankClient(host *string, port *uint) (bpb.BankServiceClient, *grpc.ClientConn, error) {
	// Dial the gRPC server using the provided configuration
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", *host, *port), // Format HttpPort as an int
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to BankService: %w", err)
	}

	client := bpb.NewBankServiceClient(conn)
	return client, conn, nil
}

func ValidDate(date string) bool {
	layout := time.RFC3339
	_, err := time.Parse(layout, date)
	if err != nil {
		return false
	} else {
		return true
	}
}
