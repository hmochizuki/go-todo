package domain

import "time"

type Status string

const (
	TODO     Status = "todo"
	DOING    Status = "doing"
	PENDING  Status = "pending"
	ARCHIVED Status = "archived"
	DONE     Status = "done"
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

type TodoService interface {
	GetAll() ([]Todo, error)
	Create(todo CreateTodoRequest) error
	Delete(id int) error
}
