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

func (h *StatisticsHandler) GetSummaryStatistics(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")
	summaryStatistics, err := h.service.GetSummaryStatistics(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(summaryStatistics)
}
