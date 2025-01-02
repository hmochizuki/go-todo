package domain

import "time"

type TodoStatus string

const (
	TODO     TodoStatus = "todo"
	DOING    TodoStatus = "doing"
	PENDING  TodoStatus = "pending"
	ARCHIVED TodoStatus = "archived"
	DONE     TodoStatus = "done"
)

type Todo struct {
	ID        uint
	Name      string
	UserID    string
	Status    TodoStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateTodoRequest struct {
	Name   string `json:"name"`
	UserID string `json:"userId"`
}

type UpdateTodoStatusRequest struct {
	Status TodoStatus `json:"status"`
}

type TodoService interface {
	GetAll() ([]Todo, error)
	Create(todo CreateTodoRequest) error
	UpdateTodoStatus(todoId uint, todo UpdateTodoStatusRequest) error
	Delete(id int) error
}
