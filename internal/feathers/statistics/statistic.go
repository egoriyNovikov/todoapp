package statistics

import "github.com/egoriynovikov/todoapp/internal/feathers/tasks"

type CompletedTasksByUser struct {
	UserID         string `json:"user_id"`
	TasksCompleted *int   `json:"tasks_completed"`
}

type TasksInCompletedByUser struct {
	UserID           string `json:"user_id"`
	TasksInCompleted *int   `json:"tasks_incompleted"`
}

type AllTasks struct {
	TasksCompleted   *int `json:"tasks_completed"`
	TasksInCompleted *int `json:"tasks_incompleted"`
}

type AllTasksByUser struct {
	UserID   string   `json:"user_id"`
	AllTasks AllTasks `json:"all_tasks"`
}

type AllCompletedTasks struct {
	TasksCompleted []*tasks.Task `json:"tasks_completed"`
}

type AllInCompletedTasks struct {
	TasksInCompleted []*tasks.Task `json:"tasks_incompleted"`
}

type GetAllUsersStatistics struct {
	UserID              string              `json:"user_id"`
	AllTasks            AllTasks            `json:"all_tasks"`
	AllCompletedTasks   AllCompletedTasks   `json:"all_completed_tasks"`
	AllInCompletedTasks AllInCompletedTasks `json:"all_incompleted_tasks"`
}
