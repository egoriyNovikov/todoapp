package tasks_repository

import (
	"context"
	"time"

	"github.com/egoriynovikov/todoapp/internal/featchers/tasks"
	"github.com/jackc/pgx/v5"
)

type postgresTaskRepository struct {
	db *pgx.Conn
}

type TaskRepository interface {
	FindByID(id string) (tasks.Task, error)
	CreateTask(task *tasks.Task) (string, error)
	UpdateTask(id string, task *tasks.Task) (string, error)
	SoftDeleteTask(id string) (string, error)
	FindAll() ([]tasks.Task, error)
}

func NewPostgresRepo(db *pgx.Conn) *postgresTaskRepository {
	return &postgresTaskRepository{db: db}
}

func (r *postgresTaskRepository) FindByID(id string) (tasks.Task, error) {
	row := r.db.QueryRow(context.Background(), "SELECT * FROM todoapp.tasks WHERE id = $1", id)
	var task tasks.Task
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.UpdatedAt, &task.AuthorUserID, &task.DeletedAt)
	if err != nil {
		return tasks.Task{}, err
	}
	return task, nil
}

func (r *postgresTaskRepository) CreateTask(task *tasks.Task) (string, error) {
	row := r.db.QueryRow(context.Background(), "INSERT INTO todoapp.tasks (title, description, completed, author_user_id) VALUES ($1, $2, $3, $4) RETURNING id", task.Title, task.Description, task.Completed, task.AuthorUserID)
	var id string
	err := row.Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *postgresTaskRepository) UpdateTask(id string, task *tasks.Task) (string, error) {
	currentTask, err := r.FindByID(id)
	if err != nil {
		return "", err
	}
	if task.Title == "" {
		task.Title = currentTask.Title
	}
	if task.Description == "" {
		task.Description = currentTask.Description
	}
	if task.Completed == false {
		task.Completed = currentTask.Completed
	}
	if task.AuthorUserID == "" {
		task.AuthorUserID = currentTask.AuthorUserID
	}
	rows, err := r.db.Query(context.Background(), "UPDATE todoapp.tasks SET title = $1, description = $2, completed = $3 WHERE id = $4 RETURNING id", task.Title, task.Description, task.Completed, id)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	return "Task updated successfully " + id, nil
}

func (r *postgresTaskRepository) SoftDeleteTask(id string) (string, error) {
	rows, err := r.db.Query(context.Background(), "UPDATE todoapp.tasks SET deleted_at = $1 WHERE id = $2 RETURNING id", time.Now(), id)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	return "Task deleted successfully " + id, nil
}

func (r *postgresTaskRepository) FindAll() ([]tasks.Task, error) {
	rows, err := r.db.Query(context.Background(), "SELECT * FROM todoapp.tasks WHERE deleted_at IS NULL")
	if err != nil {
		return []tasks.Task{}, err
	}
	defer rows.Close()

	var tasksList []tasks.Task
	for rows.Next() {
		var task tasks.Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.UpdatedAt, &task.AuthorUserID, &task.DeletedAt)
		if err != nil {
			return nil, err
		}
		tasksList = append(tasksList, task)
	}
	return tasksList, nil
}
