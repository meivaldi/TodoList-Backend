package repository

import (
	"backend_go/model/domain"
	"context"
	"database/sql"
)

type TodoListRepository interface {
	Save(ctx context.Context, tx *sql.Tx, todoList domain.TodoList) domain.TodoList
	Update(ctx context.Context, tx *sql.Tx, todoList domain.TodoList) domain.TodoList
	Delete(ctx context.Context, tx *sql.Tx, todoList domain.TodoList)
	FindById(ctx context.Context, tx *sql.Tx, todoListId int) (domain.TodoList, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.TodoList
}
