package service

import (
	"github.com/ikatseiko/todo-app-copy"
	"github.com/ikatseiko/todo-app-copy/pkg/repository"
)

type Autorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParsToken(token string) (int, error)
}

type TodoList interface {
	Create(userID int, list todo.TodoList) (int, error)
	GetAll(userID int) ([]todo.TodoList, error)
	GetByID(userID, listID int) (todo.TodoList, error)
	Delete(userID, listID int) error
	Update(userID, listID int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(userID, listID int, input todo.TodoItem) (int, error)
	GetAll(userID, listID int) ([]todo.TodoItem, error)
	GetByID(userID, itemID int) (todo.TodoItem, error)
	Delete(userID, itemID int) error
	Update(userID, itemID int, input todo.UpdateItemInput) error
}

type Service struct {
	Autorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(repos.Autorization),
		TodoList:     NewTodListService(repos.TodoList),
		TodoItem:     NewTodItemService(repos.TodoItem, repos.TodoList),
	}
}
