package controller

import (
	"backend_go/helper"
	"backend_go/model/web"
	"backend_go/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type TodoListControllerImpl struct {
	TodoListService service.TodoListService
}

func NewTodoListController(todoListService service.TodoListService) TodoListController {
	return &TodoListControllerImpl{
		TodoListService: todoListService,
	}
}

func (controller *TodoListControllerImpl) Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoListCreateRequest := web.TodoListCreateRequest{}
	helper.ReadFromRequestBody(request, &todoListCreateRequest)

	todoListResponse := controller.TodoListService.Create(request.Context(), todoListCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoListResponse,
	}

	helper.WriteToResponseBody(writter, webResponse)
}

func (controller *TodoListControllerImpl) Update(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoListUpdateRequest := web.TodoListUpdateRequest{}
	helper.ReadFromRequestBody(request, &todoListUpdateRequest)

	todoListId := params.ByName("todoListId")
	id, err := strconv.Atoi(todoListId)
	helper.PanicIfError(err)

	todoListUpdateRequest.Id = id
	todoListResponse := controller.TodoListService.Update(request.Context(), todoListUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoListResponse,
	}

	helper.WriteToResponseBody(writter, webResponse)
}

func (controller *TodoListControllerImpl) Delete(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoListId := params.ByName("todoListId")
	id, err := strconv.Atoi(todoListId)
	helper.PanicIfError(err)

	controller.TodoListService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writter, webResponse)
}

func (controller *TodoListControllerImpl) FindById(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoListId := params.ByName("todoListId")
	id, err := strconv.Atoi(todoListId)
	helper.PanicIfError(err)

	todoListResponse := controller.TodoListService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoListResponse,
	}

	helper.WriteToResponseBody(writter, webResponse)
}

func (controller *TodoListControllerImpl) FindAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoListResponses := controller.TodoListService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoListResponses,
	}

	helper.WriteToResponseBody(writter, webResponse)
}
