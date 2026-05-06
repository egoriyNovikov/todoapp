package user_service

import (
	"errors"
	"fmt"

	"github.com/egoriynovikov/todoapp/internal/featchers/users"
	user_repository "github.com/egoriynovikov/todoapp/internal/featchers/users/repository"
)

type UserService interface {
	GetUserByID(id string) (users.User, error)
	CreateUser(user *users.User) (string, error)
	GetAllUsers() ([]users.User, error)
	UpdateUser(id string, user *users.User) (string, error)
	SoftDeleteUser(id string) (string, error)
}

type service struct {
	repo user_repository.UserRepository
}

func NewService(r user_repository.UserRepository) UserService {
	return &service{repo: r}
}

func (s *service) GetUserByID(id string) (users.User, error) {
	data, err := s.repo.FindByID(id)
	if err != nil {
		fmt.Printf("failed to get user by id: %v", err)
		return users.User{}, errors.New("failed to get user by id: " + err.Error())
	}
	return data, nil
}

func (s *service) GetAllUsers() ([]users.User, error) {
	data, err := s.repo.FindAll()
	if err != nil {
		fmt.Printf("failed to get all users: %v", err)
		return []users.User{}, errors.New("failed to get all users: " + err.Error())
	}
	return data, nil
}

func (s *service) CreateUser(user *users.User) (string, error) {
	id, err := s.repo.CreateUser(user)
	if err != nil {
		fmt.Printf("failed to create user: %v\n", err)
		return "", errors.New("failed to create user: " + err.Error())
	}
	return id, nil
}

func (s *service) UpdateUser(id string, user *users.User) (string, error) {
	id, err := s.repo.UpdateUser(id, user)
	if err != nil {
		fmt.Printf("failed to update user: %v\n", err)
		return "", errors.New("failed to update user: " + err.Error())
	}
	return id, nil
}

func (s *service) SoftDeleteUser(id string) (string, error) {
	result, err := s.repo.SoftDeleteUser(id)
	if err != nil {
		fmt.Printf("failed to soft delete user: %v\n", err)
		return "", errors.New("failed to soft delete user: " + err.Error())
	}
	return result, nil
}
