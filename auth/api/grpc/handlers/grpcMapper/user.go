package grpcMapper

import (
	"github.com/onlineTraveling/auth/internal/user/domain"
	"github.com/onlineTraveling/auth/protobufs"
)

func TokenRequestToTokenDomain(w *protobufs.GetUserByTokenRequest) (*domain.Token, error) {

	return &domain.Token{
		Token: w.Token,
	}, nil
}
