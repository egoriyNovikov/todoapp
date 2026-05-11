package users

import "time"

type User struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SoftDeleteUserRequest struct {
	ID        string    `json:"id"`
	DeletedAt time.Time `json:"deleted_at"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
