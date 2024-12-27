package handler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/onlineTraveling/bank/api/service"
	"github.com/onlineTraveling/bank/internal/bank/domain"
)

type BankHandler struct {
	bankService *service.BankService
}

func NewBankHandler(bankService *service.BankService) *BankHandler {
	return &BankHandler{bankService: bankService}
}

func (h *BankHandler) CreateWallet(createWalletData []byte) {
	type CreateWalletRequest struct {
		UserID string `json:"user_id"`
	}
	var req CreateWalletRequest
	err := json.Unmarshal(createWalletData, &req)
	if err != nil {
		log.Printf("Failed to deserialize message: %v", err)
	}
	uid, _ := uuid.Parse(req.UserID)
	log.Printf("Creating wallet for user ID: %s", req.UserID)

	h.bankService.CreateWallet(context.Background(), &domain.Wallet{
		UserID: uid,
	})
}

func (h *BankHandler) Transfer(transferdata []byte) {
	type TransferRequest struct {
		FromWalletID string `json:"wallet_id_from"`
		ToWalletID   string `json:"wallet_id_to"`
		Amount       uint   `json:"amount"`
	}
	var req TransferRequest
	err := json.Unmarshal(transferdata, &req)
	if err != nil {
		log.Printf("Failed to deserialize message: %v", err)
	}
	uidFrom, err := uuid.Parse(req.FromWalletID)
	if err != nil {
		log.Printf("Failed to parse FromWalletID: %v", err)
		return
	}
	uidTo, err := uuid.Parse(req.ToWalletID)
	if err != nil {
		log.Printf("Failed to parse ToWalletID: %v", err)
		return
	}
	// log.Printf("Transfering money from ID: %s to %s", req.FromWalletID, req.ToWalletID)

	h.bankService.Transfer(context.Background(), &domain.BankTransaction{
		FromWallet: &domain.Wallet{ID: &uidFrom},
		ToWallet:   &domain.Wallet{ID: &uidTo},
		Amount:     req.Amount,
	})
}
