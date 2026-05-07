package tasks_service

import (
	"errors"
	"fmt"

	"github.com/egoriynovikov/todoapp/internal/featchers/tasks"
	tasks_repository "github.com/egoriynovikov/todoapp/internal/featchers/tasks/repository"
)

type TaskService interface {
	GetTaskByID(id string) (tasks.Task, error)
	CreateTask(task *tasks.Task) (string, error)
	UpdateTask(id string, task *tasks.Task) (string, error)
	SoftDeleteTask(id string) (string, error)
	GetAllTasks() ([]tasks.Task, error)
}

type service struct {
	repo tasks_repository.TaskRepository
}

func NewService(r tasks_repository.TaskRepository) TaskService {
	return &service{repo: r}
}

func (s *service) GetTaskByID(id string) (tasks.Task, error) {
	data, err := s.repo.FindByID(id)
	if err != nil {
		fmt.Printf("failed to get task by id: %v", err)
		return tasks.Task{}, errors.New("failed to get task by id: " + err.Error())
	}
	return data, nil
}

func (s *service) CreateTask(task *tasks.Task) (string, error) {
	id, err := s.repo.CreateTask(task)
	if err != nil {
		fmt.Printf("failed to create task: %v", err)
		return "", errors.New("failed to create task: " + err.Error())
	}
	return id, nil
}

func (s *service) UpdateTask(id string, task *tasks.Task) (string, error) {
	id, err := s.repo.UpdateTask(id, task)
	if err != nil {
		fmt.Printf("failed to update task: %v", err)
		return "", errors.New("failed to update task: " + err.Error())
	}
	return id, nil
}

func (s *service) SoftDeleteTask(id string) (string, error) {
	result, err := s.repo.SoftDeleteTask(id)
	if err != nil {
		fmt.Printf("failed to soft delete task: %v", err)
		return "", errors.New("failed to soft delete task: " + err.Error())
	}
	return result, nil
}

func (s *service) GetAllTasks() ([]tasks.Task, error) {
	data, err := s.repo.FindAll()
	if err != nil {
		fmt.Printf("failed to get all tasks: %v", err)
		return []tasks.Task{}, errors.New("failed to get all tasks: " + err.Error())
	}
	return data, nil
}
