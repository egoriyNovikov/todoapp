package core_router_http_tasks_api

import (
	"net/http"

	core_auth "github.com/egoriynovikov/todoapp/internal/core/auth"
	tasks_repository "github.com/egoriynovikov/todoapp/internal/feathers/tasks/repository"
	tasks_service "github.com/egoriynovikov/todoapp/internal/feathers/tasks/service"
	http_task_transport "github.com/egoriynovikov/todoapp/internal/feathers/tasks/transport/http"
	"github.com/jackc/pgx/v5"
)

func RegisterTaskRoutes(httpRouter *http.ServeMux, db *pgx.Conn) {
	taskRepo := tasks_repository.NewPostgresRepo(db)
	taskSvc := tasks_service.NewService(taskRepo)
	taskHdl := http_task_transport.NewTaskHandler(taskSvc)
	httpRouter.Handle("GET /task/{id}", core_auth.AuthenticateMiddleware(http.HandlerFunc(taskHdl.GetTask)))
	httpRouter.Handle("POST /task", core_auth.AuthenticateMiddleware(http.HandlerFunc(taskHdl.CreateTask)))
	httpRouter.Handle("PUT /task/{id}", core_auth.AuthenticateMiddleware(http.HandlerFunc(taskHdl.UpdateTask)))
	httpRouter.Handle("DELETE /task/{id}", core_auth.AuthenticateMiddleware(http.HandlerFunc(taskHdl.SoftDeleteTask)))
	httpRouter.Handle("GET /tasks", core_auth.AuthenticateMiddleware(http.HandlerFunc(taskHdl.GetAllTasks)))
}
