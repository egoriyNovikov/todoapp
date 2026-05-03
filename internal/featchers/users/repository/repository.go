package user_repository

import (
	"errors"

	"github.com/jackc/pgx/v5"
)

type postgresUserRepository struct {
	db *pgx.Conn
}

type UserRepository interface {
	FindByID(id string) (string, error)
}

func NewPostgresRepo(db *pgx.Conn) *postgresUserRepository {
	return &postgresUserRepository{db: db}
}

func (r *postgresUserRepository) FindByID(id string) (string, error) {
	if id == "0" {
		return "", errors.New("not found")
	}
	return "John Doe", nil
}
