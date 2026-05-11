package user_repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	core_auth "github.com/egoriynovikov/todoapp/internal/core/auth"
	"github.com/egoriynovikov/todoapp/internal/featchers/users"
	"github.com/jackc/pgx/v5"
)

type postgresUserRepository struct {
	db *pgx.Conn
}

type UserRepository interface {
	FindByID(id string) (users.User, error)
	CreateUser(user *users.User) (string, error)
	FindAll() ([]users.User, error)
	UpdateUser(id string, user *users.User) (string, error)
	SoftDeleteUser(id string) (string, error)
	FindByEmail(email string, password string) (users.User, error)
}

func NewPostgresRepo(db *pgx.Conn) *postgresUserRepository {
	return &postgresUserRepository{db: db}
}

func (r *postgresUserRepository) FindByID(id string) (users.User, error) {
	row := r.db.QueryRow(context.Background(), "SELECT * FROM todoapp.users WHERE id = $1", id)
	var user users.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return users.User{}, err
	}
	return user, nil
}

func (r *postgresUserRepository) CreateUser(user *users.User) (string, error) {
	hashedPassword := core_auth.HashPassword(user.Password)
	row := r.db.QueryRow(context.Background(), "INSERT INTO todoapp.users (name, email, password) VALUES ($1, $2, $3) RETURNING id", user.Name, user.Email, hashedPassword)
	var id string
	err := row.Scan(&id)
	if err != nil {
		fmt.Printf("failed to create user: %v\n", err)
		return "", errors.New("failed to create user: " + err.Error())
	}
	return id, nil

}

func (r *postgresUserRepository) FindAll() ([]users.User, error) {
	rows, err := r.db.Query(context.Background(), "SELECT * FROM todoapp.users WHERE deleted_at IS NULL")
	if err != nil {
		return []users.User{}, err
	}
	defer rows.Close()

	var usersList []users.User
	for rows.Next() {
		var user users.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		usersList = append(usersList, user)
	}
	return usersList, nil
}

func (r *postgresUserRepository) UpdateUser(id string, user *users.User) (string, error) {
	currentUser, err := r.FindByID(id)
	if err != nil {
		return "", err
	}

	if user.Email == "" {
		user.Email = currentUser.Email
	}
	if user.Name == "" {
		user.Name = currentUser.Name
	}
	if user.Password == "" {
		user.Password = currentUser.Password
	}

	var updatedID string
	err = r.db.QueryRow(context.Background(),
		"UPDATE todoapp.users SET name=$1, email=$2, password=$3, updated_at=NOW() WHERE id=$4 RETURNING id",
		user.Name, user.Email, user.Password, id,
	).Scan(&updatedID)

	if err != nil {
		return "", err
	}
	return "User updated successfully " + id, nil
}

func (r *postgresUserRepository) SoftDeleteUser(id string) (string, error) {
	rows, err := r.db.Query(context.Background(), "UPDATE todoapp.users SET deleted_at = $1 WHERE id = $2 RETURNING id", time.Now(), id)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	return "User deleted successfully " + id, nil
}

func (r *postgresUserRepository) FindByEmail(email string, password string) (users.User, error) {
	row := r.db.QueryRow(context.Background(), "SELECT * FROM todoapp.users WHERE email = $1", email)
	var user users.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return users.User{}, err
	}

	if password == "" || user.Password != core_auth.HashPassword(password) {
		return users.User{}, errors.New("invalid password")
	}
	return user, nil
}
