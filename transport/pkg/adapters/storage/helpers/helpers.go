package helpers

import (
	"fmt"
	"time"

	"github.com/onlineTraveling/transport/pkg/adapters/storage/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(host *string, port *uint) (pb.VehicleServiceClient, *grpc.ClientConn, error) {
	// Dial the gRPC server using the provided configuration
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", *host, *port), // Format HttpPort as an int
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to VehicleService: %w", err)
	}

	client := pb.NewVehicleServiceClient(conn)
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
