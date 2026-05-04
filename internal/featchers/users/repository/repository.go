package user_repository

import (
	"context"
	"errors"
	"fmt"

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
}

func NewPostgresRepo(db *pgx.Conn) *postgresUserRepository {
	return &postgresUserRepository{db: db}
}

func (r *postgresUserRepository) FindByID(id string) (users.User, error) {
	row := r.db.QueryRow(context.Background(), "SELECT name FROM todoapp.users WHERE id = $1", id)
	var user users.User
	err := row.Scan(&user)
	if err != nil {
		return users.User{}, err
	}
	return user, nil
}

func (r *postgresUserRepository) CreateUser(user *users.User) (string, error) {
	row := r.db.QueryRow(context.Background(), "INSERT INTO todoapp.users (name, email, password) VALUES ($1, $2, $3) RETURNING id", user.Name, user.Email, user.Password)
	var id string
	err := row.Scan(&id)
	if err != nil {
		fmt.Printf("failed to create user: %v\n", err)
		return "", errors.New("failed to create user: " + err.Error())
	}
	return id, nil

}

func (r *postgresUserRepository) FindAll() ([]users.User, error) {
	rows, err := r.db.Query(context.Background(), "SELECT * FROM todoapp.users")
	if err != nil {
		return []users.User{}, err
	}
	defer rows.Close()

	var usersList []users.User
	for rows.Next() {
		var user users.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		usersList = append(usersList, user)
	}
	return usersList, nil
}
