package app

import (
	"backend_go/controller"
	"backend_go/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(todoListController controller.TodoListController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/todolist", todoListController.FindAll)
	router.POST("/api/todolist", todoListController.Create)
	router.GET("/api/todolist/:todoListId", todoListController.FindById)
	router.PUT("/api/todolist/:todoListId", todoListController.Update)
	router.DELETE("/api/todolist/:todoListId", todoListController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
