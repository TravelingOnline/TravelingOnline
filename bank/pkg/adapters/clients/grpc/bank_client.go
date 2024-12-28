package grpc

// import (
// 	"context"
// 	"fmt"

// 	"github.com/onlineTraveling/bank/internal/bank/domain"
// 	"github.com/onlineTraveling/bank/pkg/adapters/clients/grpc/mappers"
// 	"github.com/onlineTraveling/bank/pkg/ports"
// 	"github.com/onlineTraveling/bank/protobufs"
// 	"google.golang.org/grpc"
// )

// type GRPCBankClient struct {
// 	ServiceRegistry ports.IServiceRegistry
// 	BankServiceName string
// }

// func NewGRPCBankClient(serviceRegistry ports.IServiceRegistry, bankServiceName string) *GRPCBankClient {
// 	return &GRPCBankClient{ServiceRegistry: serviceRegistry, BankServiceName: bankServiceName}
// }

// func (g *GRPCBankClient) CreatWallet(UserID string) (*domain.Response, error) {
// 	port, ip, err := g.ServiceRegistry.DiscoverService(g.BankServiceName)
// 	if err != nil {
// 		return nil, err
// 	}

// 	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", ip, port), grpc.WithInsecure())
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer conn.Close()

// 	// Create a new BankService client
// 	client := protobufs.NewBankServiceClient(conn)

// 	// Create a context
// 	ctx := context.Background()

// 	// Prepare the request
// 	request := &protobufs.CreateWalletRequest{
// 		UserID: UserID,
// 	}

// 	response, err := client.CreateWallet(ctx, request)
// 	if err != nil {
// 		return nil, err
// 	}
// 	domainResponse, err := mappers.CreateWalletResponseToMessageDomain(response)
// 	if err != nil {
// 		return nil, err
// 	}
// 	print(domainResponse)
// 	return domainResponse, nil
// }
