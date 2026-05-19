package statistics_repository

import (
	"context"

	"github.com/egoriynovikov/todoapp/internal/feathers/statistics"
	"github.com/jackc/pgx/v5"
)

type postgresStatisticsRepository struct {
	db *pgx.Conn
}

type StatisticsRepository interface {
	GetSummaryStatistics(userID string) (statistics.SummaryStatistics, error)
}

func NewPostgresRepo(db *pgx.Conn) *postgresStatisticsRepository {
	return &postgresStatisticsRepository{db: db}
}

func (r *postgresStatisticsRepository) GetSummaryStatistics(userID string) (statistics.SummaryStatistics, error) {
	rows, err := r.db.Query(context.Background(), `SELECT
  COUNT(*) FILTER (WHERE completed)     AS completed,
  COUNT(*) FILTER (WHERE NOT completed) AS pending,
  COUNT(*) AS total
	FROM todoapp.tasks
	WHERE deleted_at IS NULL
  AND ($1::uuid IS NULL OR author_user_id = $1)`, userID)
	if err != nil {
		return statistics.SummaryStatistics{}, err
	}
	defer rows.Close()
	if userID == "" {
		userID = "all"
	}
	var completed int = 0
	var pending int = 0
	var total int = 0
	err = rows.Scan(&completed, &pending, &total)
	if err != nil {
		return statistics.SummaryStatistics{}, err
	}
	return statistics.SummaryStatistics{
		UserID: userID,
		Counts: statistics.TaskCounts{
			Completed: completed,
			Pending:   pending,
			Total:     total,
		},
		Rate: float64(completed) / float64(total),
	}, nil
}
