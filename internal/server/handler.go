package server

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/labstack/echo/v4"
	"github.com/wolfeidau/lambda-golang-containers/internal/todosapi"
)

func Setup(awscfg config.Config, e *echo.Echo) error {

	todos := &todosAPI{}

	todosapi.RegisterHandlers(e, todos)

	return nil
}

type todosAPI struct {
	todos map[string]*todosapi.Todo
}

// List the available tasks
// (GET /todos)
func (ts *todosAPI) ListTodos(c echo.Context, params todosapi.ListTodosParams) error {

	result := make([]todosapi.Todo, 0)

	for _, v := range ts.todos {
		if checkStatus(params.Status, v.Status) {
			result = append(result, *v)
		}
	}

	return c.NoContent(http.StatusNotImplemented)
}

// Create a todo
// (POST /todos)
func (ts *todosAPI) CreateTodo(c echo.Context) error {
	return c.NoContent(http.StatusNotImplemented)
}

// Delete the todo
// (DELETE /todos/{todoId})
func (ts *todosAPI) DeleteTodo(c echo.Context, todoId int64) error {
	return c.NoContent(http.StatusNotImplemented)
}

// Update the todo
// (PUT /todos/{todoId})
func (ts *todosAPI) UpdateTodo(c echo.Context, todoId int64) error {
	return c.NoContent(http.StatusNotImplemented)
}

func checkStatus(filter *todosapi.ListTodosParamsStatus, status todosapi.TodoStatus) bool {
	if filter == nil {
		return true
	}

	return *filter == todosapi.ListTodosParamsStatus(status)
}
