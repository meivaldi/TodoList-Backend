package helper

import (
	"backend_go/model/domain"
	"backend_go/model/web"
)

func ToTodoListResponse(todoList domain.TodoList) web.TodoListResponse {
	return web.TodoListResponse{
		Id:          todoList.Id,
		Title:       todoList.Title,
		Description: todoList.Description,
		Thumbnail:   todoList.Thumbnail,
		Priority:    todoList.Priority,
	}
}

func ToTodoListResponses(todoLists []domain.TodoList) []web.TodoListResponse {
	var todoListResponses []web.TodoListResponse
	for _, todoList := range todoLists {
		todoListResponses = append(todoListResponses, ToTodoListResponse(todoList))
	}

	return todoListResponses
}
