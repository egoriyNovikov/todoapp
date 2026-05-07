package tasks

import "time"

type Task struct {
	ID           string     `json:"id"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Completed    bool       `json:"completed"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	AuthorUserID string     `json:"author_user_id"`
}

type CreateTaskRequest struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	Completed    bool   `json:"completed"`
	AuthorUserID string `json:"author_user_id"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type SoftDeleteTaskRequest struct {
	ID        string    `json:"id"`
	DeletedAt time.Time `json:"deleted_at"`
}
