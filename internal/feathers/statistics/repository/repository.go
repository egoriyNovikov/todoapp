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
	GetCompletedTasks(userID string) (statistics.CompletedTasksByUser, error)
	GetTasksInCompleted(userID string) (statistics.TasksInCompleted, error)
}

func NewPostgresRepo(db *pgx.Conn) *postgresStatisticsRepository {
	return &postgresStatisticsRepository{db: db}
}

func (r *postgresStatisticsRepository) GetCompletedTasks(userID string) (statistics.CompletedTasksByUser, error) {
	var tasksCompleted int
	err := r.db.QueryRow(context.Background(), "SELECT COUNT(*) FROM todoapp.tasks WHERE author_user_id = $1 AND completed = true", userID).Scan(&tasksCompleted)
	if err != nil {
		return statistics.CompletedTasksByUser{}, err
	}
	return statistics.CompletedTasksByUser{
		UserID:         userID,
		TasksCompleted: &tasksCompleted,
	}, nil
}

func (r *postgresStatisticsRepository) GetTasksCompleted(userID string) (statistics.TasksInCompleted, error) {
	var tasksIncompleted int
	err := r.db.QueryRow(context.Background(), "SELECT * FROM todoapp.tasks WHERE author_user_id = $1 AND completed = false", userID).Scan(&tasksIncompleted)
	if err != nil {
		return statistics.TasksInCompleted{}, err
	}
	return statistics.TasksInCompleted{
		UserID:           userID,
		TasksInCompleted: tasksIncompleted,
	}, nil
}
