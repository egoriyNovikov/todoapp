package core_router_http_tasks_api

import (
	"net/http"

	tasks_repository "github.com/egoriynovikov/todoapp/internal/featchers/tasks/repository"
	tasks_service "github.com/egoriynovikov/todoapp/internal/featchers/tasks/service"
	http_task_transport "github.com/egoriynovikov/todoapp/internal/featchers/tasks/transport/http"
	"github.com/jackc/pgx/v5"
)

func RegisterTaskRoutes(httpRouter *http.ServeMux, db *pgx.Conn) {
	taskRepo := tasks_repository.NewPostgresRepo(db)
	taskSvc := tasks_service.NewService(taskRepo)
	taskHdl := http_task_transport.NewTaskHandler(taskSvc)
	httpRouter.HandleFunc("GET /task/{id}", taskHdl.GetTask)
	httpRouter.HandleFunc("POST /task", taskHdl.CreateTask)
	httpRouter.HandleFunc("PUT /task/{id}", taskHdl.UpdateTask)
	httpRouter.HandleFunc("DELETE /task/{id}", taskHdl.SoftDeleteTask)
	httpRouter.HandleFunc("GET /tasks", taskHdl.GetAllTasks)
}
