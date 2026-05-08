package core_app

import (
	"context"
	"log"
	"net/http"

	core_config "github.com/egoriynovikov/todoapp/internal/core/config"
	core_middleware_header "github.com/egoriynovikov/todoapp/internal/core/middleware/header"
	core_router_http_tasks_api "github.com/egoriynovikov/todoapp/internal/core/router/http/tasks/api"
	core_router_http_users_api "github.com/egoriynovikov/todoapp/internal/core/router/http/users/api"
	"github.com/jackc/pgx/v5"
)

type App struct {
	router *http.ServeMux

	db *pgx.Conn
}

var config = core_config.NewConfig()

func NewApp() *App {
	db, err := pgx.Connect(context.Background(), "postgres://"+config.PostgresUser+":"+config.PostgresPassword+"@"+config.PostgresHost+":"+config.PostgresPort+"/"+config.PostgresDB)
	if err != nil {
		log.Fatal(err)
	}

	a := &App{
		router: http.NewServeMux(),
		db:     db,
	}
	a.mapRoutes()
	return a
}

func (a *App) mapRoutes() {
	core_router_http_tasks_api.RegisterTaskRoutes(a.router, a.db)
	core_router_http_users_api.RegisterUserRoutes(a.router, a.db)
}

func (a *App) Run() error {
	handler := core_middleware_header.JsonContentTypeMiddleware(a.router)
	return http.ListenAndServe(":8080", handler)
}

func (a *App) Close() {
	if a.db != nil {
		_ = a.db.Close(context.Background())
	}
}
