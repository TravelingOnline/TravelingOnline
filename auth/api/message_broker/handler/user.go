package handler

import (
	// 	"context"
	// 	"encoding/json"
	// 	"log"

	// 	"github.com/google/uuid"
	"github.com/onlineTraveling/auth/api/service"
	// 	"github.com/onlineTraveling/bank/internal/bank/domain"
	// 	"github.com/onlineTraveling/bank/pkg/adapters/storage/types"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) SendMessage(createWalletData []byte) {

}
