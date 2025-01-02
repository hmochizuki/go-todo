package main

import (
	"go-todo/internal/domain"
	"go-todo/internal/handlers"
	"go-todo/internal/storage"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := storage.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// マイグレーションの実行
	err = db.AutoMigrate(&domain.Todo{})
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

	// サーバー起動
	e.Start(":8080")
}
