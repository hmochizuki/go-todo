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
	e.GET("/todos", todoHandler.GetTodos)
	e.POST("/todo", todoHandler.CreateTodo)
	e.DELETE("/todo/:id", todoHandler.DeleteTodo)

	e.Start(":8080")
}
