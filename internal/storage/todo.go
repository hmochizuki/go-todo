package storage

import (
	"go-todo/internal/db"
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

func (s *TodoStorage) Create(todo domain.CreateTodoRequest) error {
	result := s.db.Create(&db.Todo{Name: todo.Name, UserID: todo.UserID})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *TodoStorage) UpdateTodoStatus(todoId uint, todo domain.UpdateTodoStatusRequest) error {
	result := s.db.Model(&db.Todo{}).Where("id = ?", todoId).Update("status", todo.Status)
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
