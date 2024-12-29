package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	rolePort "user_service/internal/role/port"
	"user_service/internal/user/domain"
	"user_service/internal/user/port"
)

var (
	ErrUserOnCreate           = errors.New("error on creating new user")
	ErrUserCreationValidation = errors.New("validation failed")
	ErrHashedPassword         = errors.New("failed to hashed the user password")
)

type service struct {
	Repo        port.Repo
	RoleService rolePort.Service
}

func NewService(repo port.Repo, roleService rolePort.Service) port.Service {
	return &service{
		Repo:        repo,
		RoleService: roleService,
	}
}
func (s *service) CreateUser(ctx context.Context, user domain.User) (domain.UserId, error) {

	err := domain.ValidateUserRegisteration(&user)
	if err != nil {
		// TODO should be logger service

		return 0, fmt.Errorf(ErrUserCreationValidation.Error(), err)
	}
	hashedPassword, err := domain.HashPassword(user.Password)
	if err != nil {
		//TODO logger
		return 0, fmt.Errorf(ErrHashedPassword.Error(), err)
	}
	user.SetPassword(hashedPassword)
	// TODO should assign default role if there is no role in user . role not yet implemented
	if user.RoleId == 0 {
		defaultRole, err := s.RoleService.GetRoleByName(ctx, "user")
		if err != nil {
			return 0, err
		}
		user.RoleId = defaultRole.ID
	}

	user.Email = domain.LowerCaseEmail(user.Email)

	userId, err := s.Repo.Create(ctx, user)
	if err != nil {
		log.Println("error on creating user :", err.Error())
		return 0, ErrUserOnCreate
	}

	return userId, nil

}

func (s *service) GetUserById(ctx context.Context, userId domain.UserId) (*domain.User, error) {
	user, err := s.Repo.GetById(ctx, userId)
	if err != nil {
		log.Println("error on finding user : ", err.Error())
		return nil, err
	}

	return user, nil
}
