package grpcMapper

import (
	"github.com/google/uuid"
	"github.com/onlineTraveling/bank/internal/bank/domain"
	"github.com/onlineTraveling/bank/protobufs"
)

func CreateWalletReqToWalletDomain(w *protobufs.CreateWalletRequest) (*domain.Wallet, error) {
	uid, err := uuid.Parse(w.UserID)
	if err != nil {
		return nil, err
	}
	return &domain.Wallet{
		UserID: uid,
	}, nil
}
func TransferReqToBankTransactionDomain(t *protobufs.TransferRequest) (*domain.BankTransaction, error) {
	var receiverUserUUID uuid.UUID
	senderUserUUID, err := uuid.Parse(t.SenderOwnerID)
	if err != nil {
		return nil, err
	}
	receiverUserUUID, err = uuid.Parse(t.ReceiverOwnerID)
	if err != nil {
		return nil, err
	}
	fromWl := &domain.Wallet{
		UserID: senderUserUUID,
	}
	toWl := &domain.Wallet{
		UserID: receiverUserUUID,
	}
	return &domain.BankTransaction{
		Amount:     uint(t.Amount),
		FromWallet: fromWl,
		ToWallet:   toWl,
	}, nil
}
