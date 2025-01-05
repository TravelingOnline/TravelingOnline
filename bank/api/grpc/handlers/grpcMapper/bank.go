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
func TransferReqToBankTransactionDomain(t *protobufs.TransferRequest) (*domain.BankTransactionRequest, error) { //checked
	senderUUID, err := uuid.Parse(t.SenderOwnerID)
	if err != nil {
		return nil, err
	}
	receiverUUID, err := uuid.Parse(t.ReceiverOwnerID)
	if err != nil {
		return nil, err
	}
	// fromWl := &domain.Wallet{
	// 	ID: &senderUserUUID,
	// }
	// toWl := &domain.Wallet{
	// 	ID: &receiverUserUUID,
	// }
	return &domain.BankTransactionRequest{
		Amount:     uint64(t.Amount),
		FromUserID: senderUUID,
		ToUserID:   receiverUUID,
	}, nil
}
