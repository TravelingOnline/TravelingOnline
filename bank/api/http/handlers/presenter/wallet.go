package presenter

import (
	"github.com/onlineTraveling/bank/internal/bank/domain"
)

type AddCardToWalletReq struct {
	CardNumber string `json:"card_number" validate:"required"`
}

type DepositReq struct {
	CardNumber string `json:"card_number" validate:"required"`
	Amount     uint   `json:"amount" validate:"required"`
}
type WithdrawReq struct {
	CardNumber string `json:"card_number" validate:"required"`
	Amount     uint   `json:"amount" validate:"required"`
}

type AddCardToWalletResp struct {
	Card *domain.CreditCard `json:"card"`
}

type WalletCardsResp struct {
	Cards []domain.CreditCard `json:"cards"`
}
type DepositResp struct {
	Wallet *domain.Wallet `json:"wallet"`
}
type WithdrawResp struct {
	Message string         `json:"message"`
	Wallet  *domain.Wallet `json:"wallet"`
}

type GetWalletResp struct {
	Wallet *domain.Wallet `json:"wallet"`
}

func AddCardToWalletReqToCard(c *AddCardToWalletReq) *domain.CreditCard {
	return &domain.CreditCard{
		Number: c.CardNumber,
	}
}

func CardToAddCardToWalletResp(c domain.CreditCard) AddCardToWalletResp {
	return AddCardToWalletResp{Card: &c}
}

func CardsToWalletCardsResp(cards []domain.CreditCard) WalletCardsResp {
	return WalletCardsResp{Cards: cards}
}

func DepositReqNumToCard(cardNum string) *domain.CreditCard {
	return &domain.CreditCard{
		Number: cardNum,
	}
}
func WithdrawReqNumToCard(cardNum string) *domain.CreditCard {
	return &domain.CreditCard{
		Number: cardNum,
	}
}

func WalletToDepositResp(wl domain.Wallet) DepositResp {
	return DepositResp{Wallet: &wl}
}

func WalletToWithdrawResp(wl domain.Wallet) DepositResp {
	return DepositResp{Wallet: &wl}
}

func WalletToGetWalletResp(wl domain.Wallet) GetWalletResp {
	return GetWalletResp{
		Wallet: &wl,
	}
}
