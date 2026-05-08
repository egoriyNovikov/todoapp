package core_router_http_users_api

import (
	"net/http"

	user_repository "github.com/egoriynovikov/todoapp/internal/featchers/users/repository"
	user_service "github.com/egoriynovikov/todoapp/internal/featchers/users/service"
	http_user_transport "github.com/egoriynovikov/todoapp/internal/featchers/users/transport/http"
	"github.com/jackc/pgx/v5"
)

func RegisterUserRoutes(httpRouter *http.ServeMux, db *pgx.Conn) {
	userRepo := user_repository.NewPostgresRepo(db)
	userSvc := user_service.NewService(userRepo)
	userHdl := http_user_transport.NewUserHandler(userSvc)

	httpRouter.HandleFunc("GET /user/{id}", userHdl.GetUser)
	httpRouter.HandleFunc("POST /user", userHdl.CreateUser)
	httpRouter.HandleFunc("PUT /user/{id}", userHdl.UpdateUser)
	httpRouter.HandleFunc("DELETE /user/{id}", userHdl.SoftDeleteUser)
	httpRouter.HandleFunc("GET /users", userHdl.GetAllUsers)
}
