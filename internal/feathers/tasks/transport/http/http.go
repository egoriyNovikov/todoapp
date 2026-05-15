package http_task_transport

import (
	"encoding/json"
	"net/http"

	core_error "github.com/egoriynovikov/todoapp/internal/core/error"
	core_middleware_logger "github.com/egoriynovikov/todoapp/internal/core/middleware/logger"
	"github.com/egoriynovikov/todoapp/internal/feathers/tasks"
	tasks_service "github.com/egoriynovikov/todoapp/internal/feathers/tasks/service"
)

type TaskHandle struct {
	service tasks_service.TaskService
}

func NewTaskHandler(s tasks_service.TaskService) *TaskHandle {
	return &TaskHandle{service: s}
}

func (h *TaskHandle) GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	task, err := h.service.GetTaskByID(id)
	if err != nil {
		core_middleware_logger.SetError(w, err)
		core_error.WriteError(w, http.StatusNotFound, "Task not found")
		return
	}
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandle) CreateTask(w http.ResponseWriter, r *http.Request) {
	var input tasks.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		core_middleware_logger.SetError(w, err)
		core_error.WriteError(w, http.StatusBadRequest, "Invalid input")
		return
	}
	task, err := h.service.CreateTask(&tasks.Task{Title: input.Title, Description: input.Description, Completed: input.Completed, AuthorUserID: input.AuthorUserID})
	if err != nil {
		core_middleware_logger.SetError(w, err)
		core_error.WriteError(w, http.StatusInternalServerError, "Failed to create task: "+err.Error())
		return
	}
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandle) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var input tasks.UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		core_middleware_logger.SetError(w, err)
		core_error.WriteError(w, http.StatusBadRequest, "Invalid input")
		return
	}
	id := r.PathValue("id")
	task, err := h.service.UpdateTask(id, &tasks.Task{Title: input.Title, Description: input.Description, Completed: input.Completed})
	if err != nil {
		core_middleware_logger.SetError(w, err)
		core_error.WriteError(w, http.StatusInternalServerError, "Failed to update task: "+err.Error())
		return
	}
	core_middleware_logger.SetResult(w, task)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandle) SoftDeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	task, err := h.service.SoftDeleteTask(id)
	if err != nil {
		core_middleware_logger.SetError(w, err)
		core_error.WriteError(w, http.StatusInternalServerError, "Failed to soft delete task: "+err.Error())
		return
	}
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandle) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasksData, err := h.service.GetAllTasks()
	if err != nil {
		core_middleware_logger.SetError(w, err)
		core_error.WriteError(w, http.StatusInternalServerError, "Failed to get all tasks: "+err.Error())
		return
	}
	core_middleware_logger.SetResult(w, tasksData)
	json.NewEncoder(w).Encode(tasksData)
}
