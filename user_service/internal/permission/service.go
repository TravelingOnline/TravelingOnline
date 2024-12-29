package permission

import (
	"context"
	"log"
	"user_service/internal/permission/domain"
	"user_service/internal/permission/port"
)

// TODO
type service struct {
	Repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		Repo: repo,
	}
}

func (s *service) CreatePermision(ctx context.Context, permision domain.Permision) (domain.PermisionID, error) {

	newPermision, err := s.Repo.Create(ctx, permision)
	if err != nil {
		log.Println("error on creating permision:", err.Error())
		return 0, err
	}
	return newPermision, nil
}

func (s *service) GetPermisionById(ctx context.Context, permisionId domain.PermisionID) (*domain.Permision, error) {
	permision, err := s.Repo.GetById(ctx, permisionId)
	if err != nil {
		log.Println("error on finding permision by id :", err.Error())
		return nil, err
	}
	return permision, nil

}

func (s *service) GetPermisionByName(ctx context.Context, name string) (*domain.Permision, error) {
	permision, err := s.Repo.GetByName(ctx, name)
	if err != nil {
		log.Println("error on finding permision by name :", err.Error())
		return nil, err
	}
	return permision, nil
}
