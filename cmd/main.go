package main

import (
	"go-todo/internal/handlers"
	"go-todo/internal/storage"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	e := echo.New()

	// サービスとハンドラのセットアップ
	todoStorage := storage.NewInMemoryTodoStorage()
	todoHandler := handlers.TodoHandler{Service: todoStorage}

	// ルートの登録
	e.GET("/todos", todoHandler.GetTodos)
	e.POST("/todo", todoHandler.CreateTodo)
	e.DELETE("/todo/:id", todoHandler.DeleteTodo)

	// サーバー起動
	e.Start(":8080")
}
