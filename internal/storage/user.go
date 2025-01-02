package storage

import (
	"go-todo/internal/db"
	"go-todo/internal/domain"

	"gorm.io/gorm"
)

type UserStorage struct {
	db *gorm.DB
}

func NewUserStorage(db *gorm.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (s *UserStorage) CreateUser(user domain.CreateUserRequest) error {
	result := s.db.Create(&db.User{Name: user.Name, Email: user.Email, Password: user.Password})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *UserStorage) GetUserById(id string) (domain.User, error) {
	var user domain.User
	result := s.db.First(&user, id).Where("deleted = false")
	if result.Error != nil {
		return domain.User{}, result.Error
	}
	return user, nil
}
