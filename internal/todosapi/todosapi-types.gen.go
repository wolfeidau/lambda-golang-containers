// Package todosapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package todosapi

import (
	"time"
)

const (
	Sigv4Scopes = "sigv4.Scopes"
)

// Defines values for TodoStatus.
const (
	TodoStatusDone TodoStatus = "done"

	TodoStatusWaiting TodoStatus = "waiting"

	TodoStatusWorking TodoStatus = "working"
)

// Todo defines model for Todo.
type Todo struct {
	// The todo creation date
	CreateDate time.Time `json:"create_date"`

	// The todo resolution date
	DoneDate *time.Time `json:"done_date,omitempty"`

	// The todo identifier
	Id int64 `json:"id"`

	// The todo state
	Status TodoStatus `json:"status"`

	// The todo title
	Title string `json:"title"`
}

// The todo state
type TodoStatus string

// ListTodosParams defines parameters for ListTodos.
type ListTodosParams struct {
	// Filters the tasks by their status
	Status *ListTodosParamsStatus `json:"status,omitempty"`
}

// ListTodosParamsStatus defines parameters for ListTodos.
type ListTodosParamsStatus string
