package repository

import (
	"backend_go/helper"
	"backend_go/model/domain"
	"context"
	"database/sql"
	"errors"
)

type TodoListRepositoryImpl struct {
}

func NewTodoListRepository() TodoListRepository {
	return &TodoListRepositoryImpl{}
}

func (repository *TodoListRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, todoList domain.TodoList) domain.TodoList {
	statement := "INSERT INTO todolist(title, description, thumbnail, priority) VALUES(?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, statement, todoList.Title, todoList.Description, todoList.Thumbnail, todoList.Priority)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	todoList.Id = int(id)
	return todoList
}

func (repository *TodoListRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, todoList domain.TodoList) domain.TodoList {
	statement := "UPDATE todolist SET title=?, description=?, thumbnail=?, priority=? WHERE id=?"
	_, err := tx.ExecContext(ctx, statement, todoList.Title, todoList.Description, todoList.Thumbnail, todoList.Priority, todoList.Id)
	helper.PanicIfError(err)

	return todoList
}

func (repository *TodoListRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, todoList domain.TodoList) {
	statement := "DELETE FROM todolist WHERE id=?"
	_, err := tx.ExecContext(ctx, statement, todoList.Id)
	helper.PanicIfError(err)
}

func (repository *TodoListRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, todoListId int) (domain.TodoList, error) {
	statement := "SELECT id, title, description, thumbnail, priority FROM todolist WHERE id=?"
	rows, err := tx.QueryContext(ctx, statement, todoListId)
	helper.PanicIfError(err)
	defer rows.Close()

	todoList := domain.TodoList{}
	if rows.Next() {
		err := rows.Scan(&todoList.Id, &todoList.Title, &todoList.Description, &todoList.Thumbnail, &todoList.Priority)
		helper.PanicIfError(err)
		return todoList, nil
	} else {
		return todoList, errors.New("TodoList is Not Found")
	}
}

func (repository *TodoListRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.TodoList {
	statement := "SELECT id, title, description, thumbnail, priority FROM todolist"
	rows, err := tx.QueryContext(ctx, statement)
	helper.PanicIfError(err)
	defer rows.Close()

	var todoLists []domain.TodoList
	for rows.Next() {
		todoList := domain.TodoList{}
		err := rows.Scan(&todoList.Id, &todoList.Title, &todoList.Description, &todoList.Thumbnail, &todoList.Priority)
		helper.PanicIfError(err)
		todoLists = append(todoLists, todoList)
	}

	return todoLists
}
