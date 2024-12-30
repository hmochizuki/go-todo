package domain

type Todo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TodoService interface {
	GetAll() ([]Todo, error)
	Add(todo Todo) (Todo, error)
	Delete(id int) error
}
