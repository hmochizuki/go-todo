package main

import (
	"go-todo/internal/db"
	"go-todo/internal/handlers"
	"go-todo/internal/storage"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func initDB(gdb *gorm.DB) error {
	err := gdb.AutoMigrate(&db.User{})
	if err != nil {
		return err
	}

	err = gdb.AutoMigrate(&db.Todo{})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// DB接続とマイグレーションの実行
	connectedDB, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	err = initDB(connectedDB)
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}

	e := echo.New()
	// TODOルートの登録
	todoStorage := storage.NewTodoStorage(connectedDB)
	todoHandler := handlers.TodoHandler{Service: todoStorage}
	e.GET("/todos", todoHandler.GetAllTodos)
	e.POST("/todo", todoHandler.CreateTodo)
	e.POST("/todo/:id/status", todoHandler.UpdateTodoStatus)
	e.DELETE("/todo/:id", todoHandler.DeleteTodo)

	// Userルートの登録
	userStorage := storage.NewUserStorage(connectedDB)
	userHandler := handlers.UserHandler{Service: userStorage}
	e.POST("/user", userHandler.CreateUser)

	e.Start(":8080")
}
