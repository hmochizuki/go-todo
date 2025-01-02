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
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateTodoRequest struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	UserID string `json:"userId"`
}

type TodoService interface {
	GetAll() ([]Todo, error)
	Create(todo CreateTodoRequest) error
	Delete(id int) error
}
