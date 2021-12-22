package service

import (
	"backend_go/model/web"
	"context"
)

type TodoListService interface {
	Create(ctx context.Context, request web.TodoListCreateRequest) web.TodoListResponse
	Update(ctx context.Context, request web.TodoListUpdateRequest) web.TodoListResponse
	Delete(ctx context.Context, todoListId int)
	FindById(ctx context.Context, todoListId int) web.TodoListResponse
	FindAll(ctx context.Context) []web.TodoListResponse
}
