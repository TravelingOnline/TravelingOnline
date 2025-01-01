package handlers

import (
	"context"
	"errors"

	"github.com/onlineTraveling/auth/api/grpc/handlers/grpcMapper"
	"github.com/onlineTraveling/auth/api/service"
	"github.com/onlineTraveling/auth/internal/user/domain"
	"github.com/onlineTraveling/auth/protobufs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

type GRPCUserHandler struct {
	protobufs.UnimplementedAuthServiceServer
	userService *service.UserService
}

func NewGRPCUserHandler(userService *service.UserService) *GRPCUserHandler {
	return &GRPCUserHandler{userService: userService}
}

func (g *GRPCUserHandler) GetUserByToken(ctx context.Context, wl *protobufs.GetUserByTokenRequest) (*protobufs.GetUserByTokenResponse, error) {
	println(11111111111)
	domaintoken, err := grpcMapper.TokenRequestToTokenDomain(wl)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s", err.Error())
	}
	print("***here*******123\n")
	res, err := g.userService.GetUserIDFromToken(ctx, domaintoken.Token)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidUserToken) {
			print(111111111)
			return nil, status.Errorf(codes.InvalidArgument, "invalid token")
		}
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	return &protobufs.GetUserByTokenResponse{UserId: res.UserId, IsAdmin: res.IsAdmin}, nil
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
