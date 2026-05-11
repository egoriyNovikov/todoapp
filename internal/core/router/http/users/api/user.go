package core_router_http_users_api

import (
	"net/http"

	core_auth "github.com/egoriynovikov/todoapp/internal/core/auth"
	user_repository "github.com/egoriynovikov/todoapp/internal/featchers/users/repository"
	user_service "github.com/egoriynovikov/todoapp/internal/featchers/users/service"
	http_user_transport "github.com/egoriynovikov/todoapp/internal/featchers/users/transport/http"
	"github.com/jackc/pgx/v5"
)

func RegisterUserRoutes(httpRouter *http.ServeMux, db *pgx.Conn) {
	userRepo := user_repository.NewPostgresRepo(db)
	userSvc := user_service.NewService(userRepo)
	userHdl := http_user_transport.NewUserHandler(userSvc)

	httpRouter.HandleFunc("POST /user/login", userHdl.Login)
	httpRouter.HandleFunc("POST /user", userHdl.CreateUser)

	httpRouter.Handle("GET /user/{id}", core_auth.AuthenticateMiddleware(http.HandlerFunc(userHdl.GetUser)))
	httpRouter.Handle("PUT /user/{id}", core_auth.AuthenticateMiddleware(http.HandlerFunc(userHdl.UpdateUser)))
	httpRouter.Handle("DELETE /user/{id}", core_auth.AuthenticateMiddleware(http.HandlerFunc(userHdl.SoftDeleteUser)))
	httpRouter.Handle("GET /users", core_auth.AuthenticateMiddleware(http.HandlerFunc(userHdl.GetAllUsers)))

}
