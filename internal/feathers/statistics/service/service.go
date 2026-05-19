package statistics_service

import (
	"github.com/egoriynovikov/todoapp/internal/feathers/statistics"
	statistics_repository "github.com/egoriynovikov/todoapp/internal/feathers/statistics/repository"
)

type StatisticsService interface {
	GetSummaryStatistics(userID string) (statistics.SummaryStatistics, error)
}

type service struct {
	repo statistics_repository.StatisticsRepository
}

func NewService(r statistics_repository.StatisticsRepository) StatisticsService {
	return &service{repo: r}
}

func (s *service) GetSummaryStatistics(userID string) (statistics.SummaryStatistics, error) {
	return s.repo.GetSummaryStatistics(userID)
}
