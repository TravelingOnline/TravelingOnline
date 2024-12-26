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
		FromUserID string `json:"user_id_from"`
		ToUserID   string `json:"user_id_to"`
		Amount     uint   `json:"amount"`
	}
	var req TransferRequest
	err := json.Unmarshal(transferdata, &req)
	if err != nil {
		log.Printf("Failed to deserialize message: %v", err)
	}
	uidFrom, err := uuid.Parse(req.FromUserID)
	if err != nil {
		log.Printf("Failed to parse FromUserID: %v", err)
		return
	}
	uidTo, err := uuid.Parse(req.ToUserID)
	if err != nil {
		log.Printf("Failed to parse ToUserID: %v", err)
		return
	}
	log.Printf("Transfering money from ID: %s to %s", req.FromUserID, req.ToUserID)

	h.bankService.Transfer(context.Background(), &domain.BankTransaction{
		FromWallet: &domain.Wallet{UserID: uidFrom},
		ToWallet:   &domain.Wallet{UserID: uidTo},
		Amount:     req.Amount,
	})
}
