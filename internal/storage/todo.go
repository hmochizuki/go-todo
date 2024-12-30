package storage

import (
	"errors"
	"go-todo/internal/domain"
	"sync"
)

type InMemoryTodoStorage struct {
	todos  []domain.Todo
	nextID int
	mutex  sync.Mutex
}

func NewInMemoryTodoStorage() *InMemoryTodoStorage {
	return &InMemoryTodoStorage{
		todos:  []domain.Todo{},
		nextID: 1,
	}
}

func (s *InMemoryTodoStorage) GetAll() ([]domain.Todo, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.todos, nil
}

func (s *InMemoryTodoStorage) Add(todo domain.Todo) (domain.Todo, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	todo.ID = s.nextID
	s.nextID++
	s.todos = append(s.todos, todo)
	return todo, nil
}

func (s *InMemoryTodoStorage) Delete(id int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for i, t := range s.todos {
		if t.ID == id {
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}
