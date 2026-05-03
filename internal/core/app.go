package core_app

import (
	"context"
	"log"
	"net/http"

	core_config "github.com/egoriynovikov/todoapp/internal/core/config"
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
	defer db.Close(context.Background())
	return a
}

func (a *App) mapRoutes() {
	userRepo := user_repository.NewPostgresRepo(a.db)
	userSvc := user_service.NewService(userRepo)
	userHdl := http_user_transport.NewUserHandler(userSvc)

	a.router.HandleFunc("GET /user/{id}", userHdl.GetUser)
}

func (a *App) Run() error {
	return http.ListenAndServe(":8080", a.router)
}
