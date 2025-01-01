package domain

import (
	"github.com/google/uuid"
)

type WalletCreditCard struct {
	WalletID     uuid.UUID
	CreditCardID uuid.UUID
}
