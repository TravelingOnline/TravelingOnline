package company

import (
	"context"

	"github.com/onlineTraveling/transport/api/pb"
	"github.com/onlineTraveling/transport/api/service"
	"github.com/onlineTraveling/transport/internal/company/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCTransportHandler struct {
	pb.UnimplementedCompanyServiceServer
	transportService *service.CompanyService
}

func NewGRPCTransportHandler(trasnsportService service.CompanyService) *GRPCTransportHandler {
	return &GRPCTransportHandler{transportService: &trasnsportService}
}

func (g *GRPCTransportHandler) CreateCompany(ctx context.Context, req *pb.CreateCompanyRequest) (*pb.CreateCompanyResponse, error) {
	domainRequest, err := PBCompanyRequest2DomainCompany(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	vID, err := g.transportService.CreateCompany(ctx, &domainRequest)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.CreateCompanyResponse{
		Id: string(*vID),
	}, nil
}

func (g *GRPCTransportHandler) UpdateCompany(ctx context.Context, req *pb.UpdateCompanyRequest) (*pb.UpdateCompanyResponse, error) {
	domainRequest, err := PBCompanyRequest2DomainCompany(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	vID, err := g.transportService.UpdateCompany(ctx, &domainRequest)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.UpdateCompanyResponse{
		Id: string(*vID),
	}, nil
}

func (g *GRPCTransportHandler) DeleteCompany(ctx context.Context, delReq *pb.DeleteCompanyRequest) (*pb.DeleteCompanyResponse, error) {
	vID := domain.CompanyID(delReq.Id)
	deletedVehcileID, err := g.transportService.DeleteCompany(ctx, &vID)
	if err != nil {
		return deletedVehcileID, status.Errorf(codes.Internal, err.Error())
	}
	return deletedVehcileID, nil
}

func (g *GRPCTransportHandler) GetByIDCompany(ctx context.Context, CompanyReq *pb.GetByIDCompanyRequest) (*pb.GetByIDCompanyResponse, error) {
	vID := domain.CompanyID(CompanyReq.Id)
	company, err := g.transportService.GetByIDCompany(ctx, &vID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	domainCompany, err := DomainCompany2PBCompanyResponse(*company)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	return domainCompany, nil
}
