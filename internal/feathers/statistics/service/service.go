package statistics_service

import (
	"github.com/egoriynovikov/todoapp/internal/feathers/statistics"
	statistics_repository "github.com/egoriynovikov/todoapp/internal/feathers/statistics/repository"
)

type StatisticsService interface {
	GetCompletedTasks(userID string) (statistics.CompletedTasksByUser, error)
	GetTasksInCompleted(userID string) (statistics.TasksInCompletedByUser, error)
	GetAllTasks() (statistics.AllTasks, error)
	GetAllTasksByUser(userID string) (statistics.AllTasksByUser, error)
	GetAllCompletedTasks(userID string) (statistics.AllCompletedTasks, error)
	GetAllInCompletedTasks(userID string) (statistics.AllInCompletedTasks, error)
	GetAllUsersStatistics(userID string) (statistics.GetAllUsersStatistics, error)
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

func (s *service) GetTasksInCompleted(userID string) (statistics.TasksInCompletedByUser, error) {
	return s.repo.GetTasksInCompleted(userID)
}

func (s *service) GetAllTasks() (statistics.AllTasks, error) {
	return s.repo.GetAllTasks()
}

func (s *service) GetAllTasksByUser(userID string) (statistics.AllTasksByUser, error) {
	return s.repo.GetAllTasksByUser(userID)
}

func (s *service) GetAllCompletedTasks(userID string) (statistics.AllCompletedTasks, error) {
	return s.repo.GetAllCompletedTasks(userID)
}

func (s *service) GetAllInCompletedTasks(userID string) (statistics.AllInCompletedTasks, error) {
	return s.repo.GetAllInCompletedTasks(userID)
}

func (s *service) GetAllUsersStatistics(userID string) (statistics.GetAllUsersStatistics, error) {
	return s.repo.GetAllUsersStatistics(userID)
}
