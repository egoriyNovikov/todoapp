package core_router_http_statistics_api

import (
	"net/http"

	core_auth "github.com/egoriynovikov/todoapp/internal/core/auth"
	statistics_repository "github.com/egoriynovikov/todoapp/internal/feathers/statistics/repository"
	statistics_service "github.com/egoriynovikov/todoapp/internal/feathers/statistics/service"
	http_statistics_transport "github.com/egoriynovikov/todoapp/internal/feathers/statistics/tranport/http"
	"github.com/jackc/pgx/v5"
)

func RegisterStatisticsRoutes(httpRouter *http.ServeMux, db *pgx.Conn) {
	statisticsRepo := statistics_repository.NewPostgresRepo(db)
	statisticsSvc := statistics_service.NewService(statisticsRepo)
	statisticsHdl := http_statistics_transport.NewStatisticsHandler(statisticsSvc)
	httpRouter.Handle("GET /statistics/summary/{userID}", core_auth.AuthenticateMiddleware(http.HandlerFunc(statisticsHdl.GetSummaryStatistics)))
}
