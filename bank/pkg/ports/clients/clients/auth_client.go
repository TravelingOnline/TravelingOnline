package clients

import "github.com/onlineTraveling/bank/internal/user/domain"

type IAuthClient interface {
	GetUserByToken(string) (*domain.User, error)
}
