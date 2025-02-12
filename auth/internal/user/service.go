package user

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/onlineTraveling/auth/internal/user/domain"
	"github.com/onlineTraveling/auth/internal/user/port"

	"github.com/onlineTraveling/auth/pkg/jwt"
	"github.com/onlineTraveling/auth/pkg/logger"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserOnCreate           = errors.New("error on creating new user")
	ErrUserCreationValidation = errors.New("validation failed")
	ErrUserNotFound           = errors.New("user not found")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateUser(ctx context.Context, user domain.User) (domain.UserID, error) {
	if err := user.Validate(); err != nil {
		return domain.UserID(uuid.Nil), fmt.Errorf("%w %w", ErrUserCreationValidation, err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error while hashing password : ", err.Error())
		return domain.UserID(uuid.Nil), ErrUserOnCreate
	}
	user.Password = string(hashedPassword)
	userID, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		log.Println("error on creating new user : ", err.Error())
		return domain.UserID(uuid.Nil), ErrUserOnCreate
	}

	return userID, nil
}
func (s *service) GetUserIDFromToken(ctx context.Context, Token string) (*jwt.UserClaims, error) {
	print("***here  555555\n")
	user, err := s.repo.GetUserIDByToken(ctx, Token)
	if err != nil {
		return nil, err
	}
	if user == nil || user.ID == "" {
		return nil, ErrUserNotFound
	}

	return user, nil

}

func (s *service) GetUserByID(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil || user.ID == uuid.Nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (s *service) GetUserByEmail(ctx context.Context, email domain.Email) (*domain.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil || user.ID == uuid.Nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (s *service) UpdateUser(ctx context.Context, user domain.User) error {
	err := s.repo.UpdateUser(ctx, user)
	if err != nil {
		logger.Error("error in update user", nil)
		return err
	}
	return nil
}

func (s *service) DeleteByID(ctx context.Context, userID domain.UserID) error {
	err := s.repo.DeleteByID(ctx, userID)
	if err != nil {
		logger.Error("can not delete user", nil)
		return err
	}
	return nil
}

func (s *service) GetUserByFilter(ctx context.Context, filter *domain.UserFilter) (*domain.User, error) {
	user, err := s.repo.GetUserByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}
