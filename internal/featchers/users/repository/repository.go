package user_repository

import "errors"

type PostgresRepo struct {
	// тут может быть *sql.DB (подключение к базе)
}

type UserRepository interface {
	FindByID(id string) (string, error)
}

func NewPostgresRepo() *PostgresRepo {
	return &PostgresRepo{}
}

func (r *PostgresRepo) FindByID(id string) (string, error) {
	if id == "0" {
		return "", errors.New("not found")
	}
	return "John Doe", nil
}
