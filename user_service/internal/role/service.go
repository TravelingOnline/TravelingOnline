package role

import (
	"context"
	"errors"
	"user_service/internal/role/domain"
	"user_service/internal/role/port"
)

var (
	ErrEmptyNameField = errors.New("the name field should not be empty")
)

type service struct {
	Repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		Repo: repo,
	}

}

func (s *service) CreateRole(ctx context.Context, role *domain.Role) (domain.RoleId, error) {
	if role.Name == "" {
		return 0, ErrEmptyNameField
	}
	roleId, err := s.Repo.Create(ctx, *role)
	if err != nil {
		return 0, err
	}
	return roleId, nil

}
func (s *service) GetRoleById(ctx context.Context, roleId domain.RoleId) (*domain.Role, error) {
	var role *domain.Role

	role, err := s.Repo.GetById(ctx, roleId)
	if err != nil {
		return nil, err
	}

	return role, nil
}
func (s *service) GetRoleByName(ctx context.Context, roleName string) (*domain.Role, error) {
	var role *domain.Role
	role, err := s.Repo.GetByName(ctx, roleName)
	if err != nil {
		return nil, err
	}
	return role, nil
}
