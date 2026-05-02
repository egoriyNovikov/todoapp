package http_user_transport

import (
	"fmt"
	"net/http"

	user_service "github.com/egoriynovikov/todoapp/internal/featchers/users/service"
)

type UserHandle struct {
	service user_service.UserService
}

func NewUserHandler(s user_service.UserService) *UserHandle {
	return &UserHandle{service: s}
}

func (h *UserHandle) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	date, err := h.service.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "User: %s", date)
}
