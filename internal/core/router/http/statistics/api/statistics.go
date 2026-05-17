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
	httpRouter.Handle("GET /statistics/completed-tasks/{userID}", core_auth.AuthenticateMiddleware(http.HandlerFunc(statisticsHdl.GetCompletedTasks)))
	httpRouter.Handle("GET /statistics/incompleted-tasks/{userID}", core_auth.AuthenticateMiddleware(http.HandlerFunc(statisticsHdl.GetTasksInCompleted)))
	httpRouter.Handle("GET /statistics/all-tasks", core_auth.AuthenticateMiddleware(http.HandlerFunc(statisticsHdl.GetAllTasks)))
	httpRouter.Handle("GET /statistics/all-tasks/{userID}", core_auth.AuthenticateMiddleware(http.HandlerFunc(statisticsHdl.GetAllTasksByUser)))
	httpRouter.Handle("GET /statistics/all-completed-tasks/{userID}", core_auth.AuthenticateMiddleware(http.HandlerFunc(statisticsHdl.GetAllCompletedTasks)))
	httpRouter.Handle("GET /statistics/all-incompleted-tasks/{userID}", core_auth.AuthenticateMiddleware(http.HandlerFunc(statisticsHdl.GetAllInCompletedTasks)))
	httpRouter.Handle("GET /statistics/all-users-statistics/{userID}", core_auth.AuthenticateMiddleware(http.HandlerFunc(statisticsHdl.GetAllUsersStatistics)))
}
