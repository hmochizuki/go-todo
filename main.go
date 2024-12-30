package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo のインスタンス作成
	e := echo.New()

	// ルートハンドラー
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// サーバーの起動
	e.Start(":8080")
}
