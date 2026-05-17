package statistics_repository

import (
	"context"

	"github.com/egoriynovikov/todoapp/internal/feathers/statistics"
	"github.com/egoriynovikov/todoapp/internal/feathers/tasks"
	"github.com/jackc/pgx/v5"
)

type postgresStatisticsRepository struct {
	db *pgx.Conn
}

type StatisticsRepository interface {
	GetCompletedTasks(userID string) (statistics.CompletedTasksByUser, error)
	GetTasksInCompleted(userID string) (statistics.TasksInCompletedByUser, error)
	GetAllTasks() (statistics.AllTasks, error)
	GetAllTasksByUser(userID string) (statistics.AllTasksByUser, error)
	GetAllCompletedTasks(userID string) (statistics.AllCompletedTasks, error)
	GetAllInCompletedTasks(userID string) (statistics.AllInCompletedTasks, error)
	GetAllUsersStatistics(userID string) (statistics.GetAllUsersStatistics, error)
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

func (r *postgresStatisticsRepository) GetTasksInCompleted(userID string) (statistics.TasksInCompletedByUser, error) {
	var tasksCompleted int
	err := r.db.QueryRow(context.Background(), "SELECT COUNT(*) FROM todoapp.tasks WHERE author_user_id = $1 AND completed = false", userID).Scan(&tasksCompleted)
	if err != nil {
		return statistics.TasksInCompletedByUser{}, err
	}
	return statistics.TasksInCompletedByUser{
		UserID:           userID,
		TasksInCompleted: &tasksCompleted,
	}, nil
}

func (r *postgresStatisticsRepository) GetAllTasks() (statistics.AllTasks, error) {
	var tasksCompleted int
	var tasksIncompleted int
	err := r.db.QueryRow(context.Background(), "SELECT COUNT(*) FROM todoapp.tasks WHERE deleted_at IS NULL AND completed = true").Scan(&tasksCompleted)
	if err != nil {
		return statistics.AllTasks{}, err
	}
	err = r.db.QueryRow(context.Background(), "SELECT COUNT(*) FROM todoapp.tasks WHERE deleted_at IS NULL AND completed = false").Scan(&tasksIncompleted)
	if err != nil {
		return statistics.AllTasks{}, err
	}
	return statistics.AllTasks{
		TasksCompleted:   &tasksCompleted,
		TasksInCompleted: &tasksIncompleted,
	}, nil
}

func (r *postgresStatisticsRepository) GetAllTasksByUser(userID string) (statistics.AllTasksByUser, error) {
	var tasksCompleted int
	var tasksIncompleted int
	err := r.db.QueryRow(context.Background(), "SELECT COUNT(*) FROM todoapp.tasks WHERE author_user_id = $1 AND deleted_at IS NULL AND completed = true", userID).Scan(&tasksCompleted)
	if err != nil {
		return statistics.AllTasksByUser{}, err
	}
	err = r.db.QueryRow(context.Background(), "SELECT COUNT(*) FROM todoapp.tasks WHERE author_user_id = $1 AND deleted_at IS NULL AND completed = false", userID).Scan(&tasksIncompleted)
	if err != nil {
		return statistics.AllTasksByUser{}, err
	}
	return statistics.AllTasksByUser{
		UserID: userID,
		AllTasks: statistics.AllTasks{
			TasksCompleted:   &tasksCompleted,
			TasksInCompleted: &tasksIncompleted,
		},
	}, nil
}

func (r *postgresStatisticsRepository) GetAllCompletedTasks(userID string) (statistics.AllCompletedTasks, error) {
	var tasksCompleted []*tasks.Task
	rows, err := r.db.Query(context.Background(), "SELECT * FROM todoapp.tasks WHERE author_user_id = $1 AND deleted_at IS NULL AND completed = true", userID)
	if err != nil {
		return statistics.AllCompletedTasks{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var task tasks.Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.UpdatedAt, &task.AuthorUserID, &task.DeletedAt)
		if err != nil {
			return statistics.AllCompletedTasks{}, err
		}
		tasksCompleted = append(tasksCompleted, &task)
	}
	return statistics.AllCompletedTasks{
		TasksCompleted: tasksCompleted,
	}, nil
}

func (r *postgresStatisticsRepository) GetAllInCompletedTasks(userID string) (statistics.AllInCompletedTasks, error) {
	var tasksIncompleted []*tasks.Task
	rows, err := r.db.Query(context.Background(), "SELECT * FROM todoapp.tasks WHERE author_user_id = $1 AND deleted_at IS NULL AND completed = false", userID)
	if err != nil {
		return statistics.AllInCompletedTasks{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var task tasks.Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.UpdatedAt, &task.AuthorUserID, &task.DeletedAt)
		if err != nil {
			return statistics.AllInCompletedTasks{}, err
		}
		tasksIncompleted = append(tasksIncompleted, &task)
	}
	return statistics.AllInCompletedTasks{
		TasksInCompleted: tasksIncompleted,
	}, nil
}

func (r *postgresStatisticsRepository) GetAllUsersStatistics(userID string) (statistics.GetAllUsersStatistics, error) {
	completedTasks, err := r.GetAllCompletedTasks(userID)
	if err != nil {
		return statistics.GetAllUsersStatistics{}, err
	}
	incompletedTasks, err := r.GetAllInCompletedTasks(userID)
	if err != nil {
		return statistics.GetAllUsersStatistics{}, err
	}
	allTasks, err := r.GetAllTasks()
	if err != nil {
		return statistics.GetAllUsersStatistics{}, err
	}
	return statistics.GetAllUsersStatistics{
		UserID:              userID,
		AllTasks:            allTasks,
		AllCompletedTasks:   completedTasks,
		AllInCompletedTasks: incompletedTasks,
	}, nil
}
