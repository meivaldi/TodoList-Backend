package test

import (
	"backend_go/app"
	"backend_go/controller"
	"backend_go/helper"
	"backend_go/middleware"
	"backend_go/repository"
	"backend_go/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/latihan_db_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validator := validator.New()

	todoListRepository := repository.NewTodoListRepository()
	todoListService := service.NewTodoListService(todoListRepository, db, validator)
	todoListController := controller.NewTodoListController(todoListService)

	router := app.NewRouter(todoListController)

	return middleware.NewAuthMiddleware(router)
}

func truncateTodoList(db *sql.DB) {
	db.Exec("TRUNCATE todolist")
}

func TestCreateTodoListSuccess(t *testing.T) {
	db := setupTestDB()
	defer db.Close()
	truncateTodoList(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(
		`{
			"title": "TodoList API",
			"desc": "TodoList API Test",
			"thumbnail": "https://meivaldi.com/test/test.png",
			"priority": 1
		}`,
	)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/todolist", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	truncateTodoList(db)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
}

func TestCreateTodoListFailed(t *testing.T) {
	db := setupTestDB()
	truncateTodoList(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(
		`{
			"title": "",
			"desc": "TodoList API Test",
			"thumbnail": "https://meivaldi.com/test/test.png",
			"priority": 1
		}`,
	)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/todolist", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
	truncateTodoList(db)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
}
