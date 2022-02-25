package service

import (
	"github.com/ikatseiko/todo-app-copy"
	"github.com/ikatseiko/todo-app-copy/pkg/repository"
)

type TodListService struct {
	repo repository.TodoList
}

func NewTodListService(repo repository.TodoList) *TodListService {
	return &TodListService{repo: repo}
}

func (s *TodListService) Create(userID int, list todo.TodoList) (int, error) {
	return s.repo.Create(userID, list)
}

func (s *TodListService) GetAll(userID int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userID)
}

func (s *TodListService) GetByID(userID, listID int) (todo.TodoList, error) {
	return s.repo.GetByID(userID, listID)
}

func (s *TodListService) Update(userID, listID int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userID, listID, input)
}

func (s *TodListService) Delete(userID, listID int) error {
	return s.repo.Delete(userID, listID)
}
