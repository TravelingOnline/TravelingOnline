package mappers

import (
	"github.com/google/uuid"
	"github.com/onlineTraveling/bank/internal/bank/domain"
	"github.com/onlineTraveling/bank/pkg/adapters/storage/types"
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
		Amount:     t.Amount,
		FromWallet: fromWl,
		ToWallet:   toWl,
	}, nil
}

// CreateWalletDomainToWalletResponse maps the Wallet domain object to CreateWalletResponse
func CreateWalletDomainToWalletResponse(wallet *domain.Wallet) *protobufs.CreateWalletResponse {
	return &protobufs.CreateWalletResponse{
		Message: "Wallet created successfully for user: " + wallet.UserID.String(),
	}
}
func CreateWalletResponseToMessageDomain(m *protobufs.CreateWalletResponse) (*domain.Response, error) {
	return &domain.Response{Message: m.Message}, nil
}

// BankTransactionDomainToTransferResponse maps the BankTransaction domain object to TransferResponse
func BankTransactionDomainToTransferResponse(transaction *domain.BankTransaction, status string) *protobufs.TransferResponse {
	return &protobufs.TransferResponse{
		SenderOwnerID:   transaction.FromWallet.UserID.String(),
		ReceiverOwnerID: transaction.ToWallet.UserID.String(),
		Amount:          uint64(transaction.Amount),
		Status:          status,
	}
}
func TransferEntitieToTransferDomain(re *protobufs.TransferResponse) (*domain.BankTransferResponse, error) {
	return &domain.BankTransferResponse{
		SenderOwnerID:   re.SenderOwnerID,
		ReceiverOwnerID: re.ReceiverOwnerID,
		Amount:          re.Amount,
		Status:          types.TransferTransactionStatus(re.Status),
	}, nil
}
