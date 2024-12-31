package mappers

import (
	"github.com/onlineTraveling/auth/internal/bank/domain"
	"github.com/onlineTraveling/auth/protobufs"
)

func CreateWalletResponseToMessageDomain(m *protobufs.CreateWalletResponse) (*domain.Response, error) {
	return &domain.Response{Message: m.Message}, nil
}
