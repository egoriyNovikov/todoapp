package statistics

type CompletedTasksByUser struct {
	UserID         string `json:"user_id"`
	TasksCompleted *int   `json:"tasks_completed"`
}

type TasksInCompleted struct {
	UserID           string `json:"user_id"`
	TasksInCompleted int    `json:"tasks_incompleted"`
}
