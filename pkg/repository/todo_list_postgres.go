package repository

import (
	"errors"
	"fmt"
	"github.com/ikatseiko/todo-app-copy"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userID int, list todo.TodoList) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, errors.New("Create | " + err.Error())
	}
	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title,description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, errors.New("Create | " + err.Error())
	}
	createUsersListsQuery := fmt.Sprintf("INSERT INTO %s (user_id,list_id) VALUES ($1, $2)", usersListsTable)

	if _, err := tx.Exec(createUsersListsQuery, userID, id); err != nil {
		tx.Rollback()
		return 0, errors.New("Create | " + err.Error())
	}
	return id, tx.Commit()
}

func (r *TodoListPostgres) GetAll(userID int) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userID)

	return lists, err
}

func (r *TodoListPostgres) GetByID(userID, listID int) (todo.TodoList, error) {
	var list todo.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description "+
		"FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND tl.id = $2",
		todoListsTable, usersListsTable)
	err := r.db.Get(&list, query, userID, listID)

	return list, err
}

func (r *TodoListPostgres) Update(userID, listID int, input todo.UpdateListInput) error {
	setValue := make([]string, 0, 2)
	args := make([]interface{}, 0, 2)
	argId := 1

	if input.Title != nil {
		setValue = append(setValue, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValue = append(setValue, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	setQuery := strings.Join(setValue, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul "+
		"WHERE tl.id = ul.list_id AND ul.user_id = $%d AND ul.list_id = $%d",
		todoListsTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, userID, listID)
	logrus.Debug("updateQuery: %s", query)
	logrus.Debug("args: %s", args)
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *TodoListPostgres) Delete(userID, listID int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul "+
		"WHERE tl.id = ul.list_id AND ul.user_id = $1 AND ul.list_id =$2",
		todoListsTable, usersListsTable)
	_, err := r.db.Exec(query, userID, listID)
	return err
}
