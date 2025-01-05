package domain

import (
	"github.com/google/uuid"
	"github.com/onlineTraveling/bank/pkg/adapters/storage/types"
)

type Wallet struct {
	ID             *uuid.UUID `json:"id"`
	IsSystemWallet bool       `json:"is_system_wallet"`
	UserID         uuid.UUID  `json:"user_id"`
	Balance        uint64     `json:"balance"`
}

type BankTransaction struct {
	Amount     uint64
	Status     types.TransferTransactionStatus
	FromWallet *Wallet
	ToWallet   *Wallet
}
type BankTransactionRequest struct {
	FromUserID uuid.UUID
	ToUserID   uuid.UUID
	Amount     uint64
	Status     types.TransferTransactionStatus
}
type BankTransferResponse struct {
	SenderOwnerID   string
	ReceiverOwnerID string
	Amount          uint64
	Status          types.TransferTransactionStatus
}
