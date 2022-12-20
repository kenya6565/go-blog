package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// グローバル変数
var e = createMux()

func main() {
	// / にアクセスした時にarticleIndex()が発火する
	e.GET("/", articleIndex)

	// Webサーバーをポート番号 8080 で起動する
	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	// アプリケーションインスタンスを生成
  e := echo.New()

  // アプリケーションに各種ミドルウェアを設定
  e.Use(middleware.Recover())
  e.Use(middleware.Logger())
  e.Use(middleware.Gzip())

  // アプリケーションインスタンスを返却
  return e
}

func articleIndex(c echo.Context) error {
	// ステータスコード 200 で、"Hello, World!" という文字列をレスポンス
	return c.String(http.StatusOK, "Hello, World!")
}
