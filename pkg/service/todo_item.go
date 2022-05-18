package service

import (
	"github.com/ikatseiko/todo-app-copy"
	"github.com/ikatseiko/todo-app-copy/pkg/repository"
)

type TodItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodItemService {
	return &TodItemService{repo: repo, listRepo: listRepo}
}

func (s *TodItemService) Create(userID, listID int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetByID(userID, listID)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(listID, item)
}

func (s *TodItemService) GetAll(userID, listID int) ([]todo.TodoItem, error) {
	_, err := s.listRepo.GetByID(userID, listID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetAll(userID, listID)
}

func (s *TodItemService) GetByID(userID, itemID int) (todo.TodoItem, error) {
	return s.repo.GetByID(userID, itemID)
}

func (s *TodItemService) Delete(userID, itemID int) error {
	return s.repo.Delete(userID, itemID)
}

func (s *TodItemService) Update(userID, itemID int, input todo.UpdateItemInput) error {
	return s.repo.Update(userID, itemID, input)
}
