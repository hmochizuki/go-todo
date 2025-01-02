package storage

import (
	"go-todo/internal/domain"

	"gorm.io/gorm"
)

type TodoStorage struct {
	todos []domain.Todo
	db    *gorm.DB
}

func NewTodoStorage(db *gorm.DB) *TodoStorage {
	return &TodoStorage{db: db}
}

func (s *TodoStorage) Create(todo domain.Todo) error {
	result := s.db.Create(&todo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *TodoStorage) GetAll() ([]domain.Todo, error) {
	result := s.db.Find(&s.todos)
	if result.Error != nil {
		return nil, result.Error
	}
	return s.todos, nil
}

func (s *TodoStorage) Delete(id int) error {
	result := s.db.Delete(&domain.Todo{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
