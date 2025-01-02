package db

import "time"

type Todo struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	UserID    string    `gorm:"foreignKey:UserID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Deleted   bool      `gorm:"default:false"`
}

type User struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Deleted   bool      `gorm:"default:false"`
}
