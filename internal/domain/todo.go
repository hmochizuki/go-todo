package domain

type Todo struct {
	ID   uint
	Name string
}

type TodoService interface {
	GetAll() ([]Todo, error)
	Create(todo Todo) error
	Delete(id int) error
}
