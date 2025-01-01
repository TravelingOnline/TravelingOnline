package grpc

import (
	"context"
	"fmt"
	"log"

	"github.com/onlineTraveling/auth/internal/user/domain"
	"github.com/onlineTraveling/auth/pkg/adapters/clients/grpc/mappers"
	"github.com/onlineTraveling/auth/pkg/ports"
	"github.com/onlineTraveling/bank/protobufs"
	"google.golang.org/grpc"
)

type GRPCBankClient struct {
	ServiceRegistry ports.IServiceRegistry
	BankServiceName string
}

func NewGRPCBankClient(serviceRegistry ports.IServiceRegistry, bankServiceName string) *GRPCBankClient {
	return &GRPCBankClient{ServiceRegistry: serviceRegistry, BankServiceName: bankServiceName}
}

func (g *GRPCBankClient) CreateWallet(UserID string) error { /////////
	port, ip, err := g.ServiceRegistry.DiscoverService(g.BankServiceName)
	if err != nil {
		return err
	}

	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", ip, port), grpc.WithInsecure())
	if err != nil {
		return err
	}

	defer conn.Close()

	// Create a new BankService client
	client := protobufs.NewBankServiceClient(conn)

	// Create a context
	ctx := context.Background()

	// Prepare the request
	request := &protobufs.CreateWalletRequest{
		UserID: UserID,
	}

	// Call the GetUserByToken method
	response, err := client.CreateWallet(ctx, request)
	if err != nil {
		return domain.ErrInvalidUserID
	}
	messageResponse, err := mappers.CreateWalletResponseToMessageDomain(response)
	if err != nil {
		return err
	}
	log.Print(messageResponse)
	return nil
}
