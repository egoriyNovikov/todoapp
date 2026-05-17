package http_statistics_transport

import (
	"encoding/json"
	"net/http"

	statistics_service "github.com/egoriynovikov/todoapp/internal/feathers/statistics/service"
)

type StatisticsHandler struct {
	service statistics_service.StatisticsService
}

func NewStatisticsHandler(s statistics_service.StatisticsService) *StatisticsHandler {
	return &StatisticsHandler{service: s}
}

func (h *StatisticsHandler) GetCompletedTasks(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")
	completedTasks, err := h.service.GetCompletedTasks(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(completedTasks)
}

func (h *StatisticsHandler) GetTasksInCompleted(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")
	incompletedTasks, err := h.service.GetTasksInCompleted(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(incompletedTasks)
}

func (h *StatisticsHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(allTasks)
}

func (h *StatisticsHandler) GetAllTasksByUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")
	allTasks, err := h.service.GetAllTasksByUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(allTasks)
}

func (h *StatisticsHandler) GetAllCompletedTasks(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")
	allCompletedTasks, err := h.service.GetAllCompletedTasks(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(allCompletedTasks)
}

func (h *StatisticsHandler) GetAllInCompletedTasks(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")
	allInCompletedTasks, err := h.service.GetAllInCompletedTasks(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(allInCompletedTasks)
}

func (h *StatisticsHandler) GetAllUsersStatistics(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")
	allUsersStatistics, err := h.service.GetAllUsersStatistics(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(allUsersStatistics)
}
