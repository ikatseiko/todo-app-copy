package service

import (
	"github.com/ikatseiko/todo-app-copy"
	"github.com/ikatseiko/todo-app-copy/pkg/repository"
)

type AuthService struct {
	repo repository.Autorization
}

func NewAuthService(repo repository.Autorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	return s.repo.CreateUser(user)
}
