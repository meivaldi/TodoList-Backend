package service

import (
	"backend_go/exception"
	"backend_go/helper"
	"backend_go/model/domain"
	"backend_go/model/web"
	"backend_go/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type TodoListServiceImpl struct {
	TodoListRepository repository.TodoListRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewTodoListService(todoListRepo repository.TodoListRepository, DB *sql.DB, Validate *validator.Validate) TodoListService {
	return &TodoListServiceImpl{
		TodoListRepository: todoListRepo,
		DB:                 DB,
		Validate:           Validate,
	}
}

func (service *TodoListServiceImpl) Create(ctx context.Context, request web.TodoListCreateRequest) web.TodoListResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todoList := domain.TodoList{
		Title:       request.Title,
		Description: request.Description,
		Thumbnail:   request.Thumbnail,
		Priority:    request.Priority,
	}

	todoList = service.TodoListRepository.Save(ctx, tx, todoList)

	return helper.ToTodoListResponse(todoList)
}

func (service *TodoListServiceImpl) Update(ctx context.Context, request web.TodoListUpdateRequest) web.TodoListResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todoList, err := service.TodoListRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	//update field
	todoList.Title = request.Title
	todoList.Description = request.Description
	todoList.Thumbnail = request.Thumbnail
	todoList.Priority = request.Priority

	todoList = service.TodoListRepository.Update(ctx, tx, todoList)

	return helper.ToTodoListResponse(todoList)
}

func (service *TodoListServiceImpl) Delete(ctx context.Context, todoListId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todoList, err := service.TodoListRepository.FindById(ctx, tx, todoListId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.TodoListRepository.Delete(ctx, tx, todoList)
}

func (service *TodoListServiceImpl) FindById(ctx context.Context, todoListId int) web.TodoListResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todoList, err := service.TodoListRepository.FindById(ctx, tx, todoListId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToTodoListResponse(todoList)
}

func (service *TodoListServiceImpl) FindAll(ctx context.Context) []web.TodoListResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todoLists := service.TodoListRepository.FindAll(ctx, tx)

	return helper.ToTodoListResponses(todoLists)
}
