package core_app

import (
	"context"
	"log"
	"net/http"

	core_config "github.com/egoriynovikov/todoapp/internal/core/config"
	core_middleware_header "github.com/egoriynovikov/todoapp/internal/core/middleware/header"
	user_repository "github.com/egoriynovikov/todoapp/internal/featchers/users/repository"
	user_service "github.com/egoriynovikov/todoapp/internal/featchers/users/service"
	http_user_transport "github.com/egoriynovikov/todoapp/internal/featchers/users/transport/http"
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
	userRepo := user_repository.NewPostgresRepo(a.db)
	userSvc := user_service.NewService(userRepo)
	userHdl := http_user_transport.NewUserHandler(userSvc)

	a.router.HandleFunc("GET /user/{id}", userHdl.GetUser)
	a.router.HandleFunc("POST /user", userHdl.CreateUser)
	a.router.HandleFunc("PUT /user/{id}", userHdl.UpdateUser)
	a.router.HandleFunc("DELETE /user/{id}", userHdl.SoftDeleteUser)
	a.router.HandleFunc("GET /users", userHdl.GetAllUsers)
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
