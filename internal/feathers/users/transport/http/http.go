package http_user_transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_auth "github.com/egoriynovikov/todoapp/internal/core/auth"
	core_middleware_logger "github.com/egoriynovikov/todoapp/internal/core/middleware/logger"
	"github.com/egoriynovikov/todoapp/internal/feathers/users"
	user_service "github.com/egoriynovikov/todoapp/internal/feathers/users/service"
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
		core_middleware_logger.SetError(w, err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	core_middleware_logger.SetResult(w, date)
	json.NewEncoder(w).Encode(date)
}

func (h *UserHandle) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input users.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		core_middleware_logger.SetError(w, err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	id, err := h.service.CreateUser(&users.User{Name: input.Name, Email: input.Email, Password: input.Password})
	if err != nil {
		core_middleware_logger.SetError(w, err)
		fmt.Fprintf(w, "failed to create user: %v", err)
		return
	}
	core_middleware_logger.SetResult(w, id)
	json.NewEncoder(w).Encode(id)
}

func (h *UserHandle) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var input users.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		core_middleware_logger.SetError(w, err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	id := r.PathValue("id")
	id, err := h.service.UpdateUser(id, &users.User{Name: input.Name, Email: input.Email, Password: input.Password})
	if err != nil {
		core_middleware_logger.SetError(w, err)
		fmt.Fprintf(w, "failed to update user: %v", err)
		return
	}
	core_middleware_logger.SetResult(w, id)
	json.NewEncoder(w).Encode(id)
}

func (h *UserHandle) SoftDeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	id, err := h.service.SoftDeleteUser(id)
	if err != nil {
		core_middleware_logger.SetError(w, err)
		fmt.Fprintf(w, "failed to soft delete user: %v", err)
		return
	}
	core_middleware_logger.SetResult(w, id)
	json.NewEncoder(w).Encode(id)
}

func (h *UserHandle) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		core_middleware_logger.SetError(w, err)
		fmt.Fprintf(w, "failed to get all users: %v", err)
		return
	}
	core_middleware_logger.SetResult(w, users)
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandle) Login(w http.ResponseWriter, r *http.Request) {
	var input users.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input ", http.StatusBadRequest)
		return
	}
	user, err := h.service.FindByEmail(input.Email, input.Password)
	if err != nil {
		core_middleware_logger.SetError(w, err)
		http.Error(w, "Failed to find user "+err.Error(), http.StatusInternalServerError)
		return
	}
	token, err := core_auth.CreateToken(user)
	if err != nil {
		core_middleware_logger.SetError(w, err)
		http.Error(w, "Failed to login "+err.Error(), http.StatusInternalServerError)
		return
	}
	core_middleware_logger.SetResult(w, token)
	json.NewEncoder(w).Encode(token)
}
