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

func (s *TodoStorage) GetById(id uint) (domain.Todo, error) {
	var todo domain.Todo
	result := s.db.First(&todo, id).Where("deleted = false")
	if result.Error != nil {
		return domain.Todo{}, result.Error
	}
	return todo, nil
}

func (s *TodoStorage) Delete(id uint) error {
	result := s.db.Model(&db.Todo{}).Where("id = ?", id).Update("deleted", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
