package repository

import (
	"github.com/ikatseiko/todo-app-copy"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (s *AuthPostgres) CreateUser(user todo.User) (int, error) {
	return 0, nil
}