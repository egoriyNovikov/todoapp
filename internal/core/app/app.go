package app

import (
	"net/http"

	user_repository "github.com/egoriynovikov/todoapp/internal/featchers/users/repository"
	user_service "github.com/egoriynovikov/todoapp/internal/featchers/users/service"
	http_user_transport "github.com/egoriynovikov/todoapp/internal/featchers/users/transport/http"
)

type App struct {
	router *http.ServeMux
}

func NewApp() *App {
	a := &App{
		router: http.NewServeMux(),
	}
	a.mapRoutes()
	return a
}

func (a *App) mapRoutes() {
	userRepo := user_repository.NewPostgresRepo()
	userSvc := user_service.NewService(userRepo)
	userHdl := http_user_transport.NewUserHandler(userSvc)

	a.router.HandleFunc("GET /user/{id}", userHdl.GetUser)
}

func (a *App) Run() error {
	return http.ListenAndServe(":8080", a.router)
}
