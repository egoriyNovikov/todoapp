package http_task_transport

import (
	"encoding/json"
	"net/http"

	"github.com/egoriynovikov/todoapp/internal/featchers/tasks"
	tasks_service "github.com/egoriynovikov/todoapp/internal/featchers/tasks/service"
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
		http.Error(w, "Task not found "+err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandle) CreateTask(w http.ResponseWriter, r *http.Request) {
	var input tasks.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input ", http.StatusBadRequest)
		return
	}
	task, err := h.service.CreateTask(&tasks.Task{Title: input.Title, Description: input.Description, Completed: input.Completed, AuthorUserID: input.AuthorUserID})
	if err != nil {
		http.Error(w, "Failed to create task "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandle) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var input tasks.UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input ", http.StatusBadRequest)
		return
	}
	id := r.PathValue("id")
	task, err := h.service.UpdateTask(id, &tasks.Task{Title: input.Title, Description: input.Description, Completed: input.Completed})
	if err != nil {
		http.Error(w, "Failed to update task "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandle) SoftDeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	task, err := h.service.SoftDeleteTask(id)
	if err != nil {
		http.Error(w, "Failed to soft delete task "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandle) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		http.Error(w, "Failed to get all tasks "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}
