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
