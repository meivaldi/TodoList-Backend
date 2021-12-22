package service

import (
	"backend_go/model/web"
	"backend_go/repository"
	"context"
	"database/sql"
)

type TodoListServiceImpl struct {
	TodoListRepository repository.TodoListRepository
	DB                 *sql.DB
}

func (service *TodoListServiceImpl) Create(ctx context.Context, request web.TodoListCreateRequest) web.TodoListResponse {
	panic("implement me")
}

func (service *TodoListServiceImpl) Update(ctx context.Context, request web.TodoListUpdateRequest) web.TodoListResponse {
	panic("implement me")
}

func (service *TodoListServiceImpl) Delete(ctx context.Context, todoListId int) {
	panic("implement me")
}

func (service *TodoListServiceImpl) FindById(ctx context.Context, todoListId int) web.TodoListResponse {
	panic("implement me")
}

func (service *TodoListServiceImpl) FindAll(ctx context.Context) []web.TodoListResponse {
	panic("implement me")
}
