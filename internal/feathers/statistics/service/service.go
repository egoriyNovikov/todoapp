package statistics_service

import (
	"github.com/egoriynovikov/todoapp/internal/feathers/statistics"
	statistics_repository "github.com/egoriynovikov/todoapp/internal/feathers/statistics/repository"
)

type StatisticsService interface {
	GetCompletedTasks(userID string) (statistics.CompletedTasksByUser, error)
	GetTasksInCompleted(userID string) (statistics.TasksInCompleted, error)
}

type service struct {
	repo statistics_repository.StatisticsRepository
}

func NewService(r statistics_repository.StatisticsRepository) StatisticsService {
	return &service{repo: r}
}

func (s *service) GetCompletedTasks(userID string) (statistics.CompletedTasksByUser, error) {
	return s.repo.GetCompletedTasks(userID)
}

func (s *service) GetTasksInCompleted(userID string) (statistics.TasksInCompleted, error) {
	return s.repo.GetTasksInCompleted(userID)
}
