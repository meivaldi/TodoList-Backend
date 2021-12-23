package main

import (
	"backend_go/app"
	"backend_go/controller"
	"backend_go/exception"
	"backend_go/helper"
	"backend_go/middleware"
	"backend_go/repository"
	"backend_go/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validator := validator.New()

	todoListRepository := repository.NewTodoListRepository()
	todoListService := service.NewTodoListService(todoListRepository, db, validator)
	todoListController := controller.NewTodoListController(todoListService)

	router := httprouter.New()
	router.GET("/api/todolist", todoListController.FindAll)
	router.POST("/api/todolist", todoListController.Create)
	router.GET("/api/todolist/:todoListId", todoListController.FindById)
	router.PUT("/api/todolist/:todoListId", todoListController.Update)
	router.DELETE("/api/todolist/:todoListId", todoListController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
