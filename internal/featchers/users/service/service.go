package user_service

import user_repository "github.com/egoriynovikov/todoapp/internal/featchers/users/repository"

type UserService interface {
	GetUserByID(id string) (string, error)
}

type service struct {
	repo user_repository.UserRepository
}

func NewService(r user_repository.UserRepository) UserService {
	return &service{repo: r}
}

func (s *service) GetUserByID(id string) (string, error) {
	return s.repo.FindByID(id)
}
