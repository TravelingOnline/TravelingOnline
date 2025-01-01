package helpers

import (
	"fmt"
	"time"

	bpb "github.com/onlineTraveling/transport/pkg/adapters/storage/bank-pb"
	vpb "github.com/onlineTraveling/transport/pkg/adapters/storage/vehicle-pb"
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
	layout := "2006-01-02"
	_, err := time.Parse(layout, date)
	return err == nil
}
