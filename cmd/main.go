package main

import (
	"go-todo/internal/handlers"
	"go-todo/internal/storage"

	"github.com/labstack/echo/v4"
)

func main() {
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
