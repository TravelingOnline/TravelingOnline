package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	codeVerficationDomain "github.com/onlineTraveling/auth/internal/codeVerification/domain"
	codeVerificationPort "github.com/onlineTraveling/auth/internal/codeVerification/port"
	"github.com/onlineTraveling/auth/internal/user"
	"github.com/onlineTraveling/auth/internal/user/domain"
	userPort "github.com/onlineTraveling/auth/internal/user/port"
	"github.com/onlineTraveling/auth/pkg/adapters/storage/types"
	"github.com/onlineTraveling/auth/pkg/helper"
	"github.com/onlineTraveling/auth/pkg/jwt"
	"github.com/onlineTraveling/auth/pkg/logger"
	helperTime "github.com/onlineTraveling/auth/pkg/time"
	"github.com/onlineTraveling/auth/protobufs"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"

	jwt2 "github.com/golang-jwt/jwt/v5"
)

var (
	ErrPasswordNotMatch = errors.New("not match password")
)

type UserService struct {
	svc                    userPort.Service
	authSecret             string
	expMin, refreshExpMin  uint
	codeVerficationServise codeVerificationPort.Service
}

func NewUserService(svc userPort.Service, authSecret string, expMin, refreshExpMin uint, codeVerificationSvc codeVerificationPort.Service) *UserService {
	return &UserService{
		svc:                    svc,
		authSecret:             authSecret,
		expMin:                 expMin,
		refreshExpMin:          refreshExpMin,
		codeVerficationServise: codeVerificationSvc,
	}
}

var (
	ErrUserCreationValidation = user.ErrUserCreationValidation
	ErrUserOnCreate           = user.ErrUserOnCreate
	ErrUserNotFound           = user.ErrUserNotFound
)

type SignUpFirstResponseWrapper struct {
	RequestTimestamp int64                              `json:"requestTimestamp"`
	Data             *protobufs.UserSignUpFirstResponse `json:"data"`
}
type SignUpSecondResponseWrapper struct {
	RequestTimestamp int64                               `json:"requestTimestamp"`
	Data             *protobufs.UserSignUpSecondResponse `json:"data"`
}

func (s *UserService) SignUp(ctx context.Context, req *protobufs.UserSignUpFirstRequest) (*SignUpFirstResponseWrapper, error) {
	userID, err := s.svc.CreateUser(ctx, domain.User{

		Email:        domain.Email(req.GetEmail()),
		PasswordHash: req.GetPassword(),
	})

	if err != nil {
		return nil, err
	}

	code := strconv.Itoa(helper.GetRandomCode())

	s.codeVerficationServise.Send(ctx, codeVerficationDomain.NewCodeVerification(userID, fmt.Sprint(code), codeVerficationDomain.CodeVerificationTypeEmail, true, time.Minute*2))

	// go helper.SendEmail(req.GetEmail())
	// go helper.SendEmail(req.GetEmail(), code)
	response := &SignUpFirstResponseWrapper{
		RequestTimestamp: time.Now().Unix(),
		Data: &protobufs.UserSignUpFirstResponse{
			UserId: uint64(userID),
		},
	}

	return response, nil
}

func (s *UserService) SignUpCodeVerification(ctx context.Context, req *protobufs.UserSignUpSecondRequest) (*SignUpSecondResponseWrapper, error) {
	_, err := s.svc.GetUserByID(ctx, domain.UserID(req.GetUserId()))
	if err != nil {
		return nil, err
	}
	ok, err := s.codeVerficationServise.CheckUserCodeVerificationValue(ctx, domain.UserID(req.GetUserId()), req.GetCode())
	if err != nil {
		return nil, err
	}
	if ok {

		accessToken, err := jwt.CreateToken([]byte(s.authSecret), &jwt.UserClaims{
			RegisteredClaims: jwt2.RegisteredClaims{
				ExpiresAt: jwt2.NewNumericDate(helperTime.AddMinutes(s.expMin, true)),
			},
			UserID: uint(req.GetUserId()),
		})
		if err != nil {
			return nil, err
		}

		refreshToken, err := jwt.CreateToken([]byte(s.authSecret), &jwt.UserClaims{
			RegisteredClaims: jwt2.RegisteredClaims{
				ExpiresAt: jwt2.NewNumericDate(helperTime.AddMinutes(s.refreshExpMin, true)),
			},
			UserID: uint(req.GetUserId()),
		})

		if err != nil {
			return nil, err
		}

		response := &SignUpSecondResponseWrapper{
			RequestTimestamp: time.Now().Unix(), // Get current UNIX timestamp
			Data: &protobufs.UserSignUpSecondResponse{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			},
		}
		return response, nil
	} else {
		return nil, nil
	}

}
func (s *UserService) SignIn(ctx context.Context, req *protobufs.UserSignInRequest) (*SignUpSecondResponseWrapper, error) {
	user, err := s.svc.GetUserByEmail(ctx, domain.Email(req.GetEmail()))
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.GetPassword()))
	if err != nil {
		return nil, ErrPasswordNotMatch
	}
	accessToken, err := jwt.CreateToken([]byte(s.authSecret), &jwt.UserClaims{
		RegisteredClaims: jwt2.RegisteredClaims{
			ExpiresAt: jwt2.NewNumericDate(helperTime.AddMinutes(s.expMin, true)),
		},
		UserID: uint(user.ID),
	})
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.CreateToken([]byte(s.authSecret), &jwt.UserClaims{
		RegisteredClaims: jwt2.RegisteredClaims{
			ExpiresAt: jwt2.NewNumericDate(helperTime.AddMinutes(s.refreshExpMin, true)),
		},
		UserID: uint(user.ID),
	})

	if err != nil {
		return nil, err
	}

	response := &SignUpSecondResponseWrapper{
		RequestTimestamp: time.Now().Unix(), // Get current UNIX timestamp
		Data: &protobufs.UserSignUpSecondResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
	return response, nil

}

func (s *UserService) GetByID(ctx context.Context, id uint) (*protobufs.User, error) {
	user, err := s.svc.GetUserByID(ctx, domain.UserID(id))
	if err != nil {
		return nil, err
	}

	return &protobufs.User{
		Id:           uint64(user.ID),
		Email:        string(user.Email),
		PasswordHash: user.PasswordHash,

		CreatedAt: timestamppb.New(user.CreatedAt),
		DeletedAt: timestamppb.New(user.DeletedAt), // Handle DeletedAt if needed
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}, nil
}

func (s *UserService) Update(ctx context.Context, user *types.User) error {
	err := s.svc.UpdateUser(ctx, domain.User{
		ID:           domain.UserID(user.ID),
		Email:        domain.Email(user.Email),
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	})
	if err != nil {
		logger.Error("update user error", nil)
		return err
	}
	return nil
}

func (s *UserService) DeleteByID(ctx context.Context, userID int) error {
	err := s.svc.DeleteByID(ctx, domain.UserID(userID))
	if err != nil {
		logger.Error("can not delete user", nil)
		return err
	}

	logger.Info("deleted user with id "+strconv.Itoa(int(userID)), nil)
	return nil
}
