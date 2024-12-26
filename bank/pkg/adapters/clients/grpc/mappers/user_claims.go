package mappers

// import (
// 	"github.com/onlineTraveling/bank/internal/user/domain"
// 	"github.com/onlineTraveling/bank/protobufs"

// 	"github.com/google/uuid"
// )

// func UserClaimsToDomain(p *protobufs.GetUserByTokenResponse) (*domain.User, error) {
// 	u, err := uuid.Parse(p.UserId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &domain.User{
// 		ID:      u,
// 		IsAdmin: p.IsAdmin,
// 	}, nil
// }
