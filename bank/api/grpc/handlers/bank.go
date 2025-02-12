package handlers

import (
	"context"
	"errors"

	"github.com/onlineTraveling/bank/api/grpc/handlers/grpcMapper"
	"github.com/onlineTraveling/bank/api/service"
	"github.com/onlineTraveling/bank/internal/bank/port"
	"github.com/onlineTraveling/bank/protobufs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

type GRPCBankHandler struct {
	protobufs.UnimplementedBankServiceServer
	bankService *service.BankService
}

func NewGRPCBankHandler(bankService *service.BankService) *GRPCBankHandler {
	return &GRPCBankHandler{bankService: bankService}
}

func (g *GRPCBankHandler) CreateWallet(ctx context.Context, wl *protobufs.CreateWalletRequest) (*protobufs.CreateWalletResponse, error) {
	domainWallet, err := grpcMapper.CreateWalletReqToWalletDomain(wl)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s", err.Error())
	}
	_, err = g.bankService.CreateWallet(ctx, domainWallet)
	if err != nil {
		if errors.Is(err, port.ErrUserAlreadyHasWallet) {
			return nil, status.Errorf(codes.AlreadyExists, "wallet already exists")
		}
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	return &protobufs.CreateWalletResponse{Message: "wallet created"}, nil
}
func (g *GRPCBankHandler) Transfer(ctx context.Context, tr *protobufs.TransferRequest) (*protobufs.TransferResponse, error) {
	domainTransaction, err := grpcMapper.TransferReqToBankTransactionDomain(tr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s", err.Error())
	}
	createdTransaction, err := g.bankService.Transfer(ctx, domainTransaction)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	receiverOwnerID := createdTransaction.ToUserID.String()
	return &protobufs.TransferResponse{
		SenderOwnerID:   createdTransaction.FromUserID.String(),
		ReceiverOwnerID: receiverOwnerID,
		Amount:          uint64(createdTransaction.Amount),
		Status:          string(createdTransaction.Status),
	}, nil
}

type HealthServer struct {
	grpc_health_v1.HealthServer
}

// Check implements Health.Check
func (s *HealthServer) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}
