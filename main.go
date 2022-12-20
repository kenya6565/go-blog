package main

import (
	"net/http"
	"time"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const templPath = "src/template/"

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
	// return c.String(http.StatusOK, "Hello, World!")

	data := map[string]interface{}{
		"Message": "Hello, World",
		"Now": time.Now(),
	}
	return render(c, "article/index.html", data)
}

func htmlBlob(file string, data map[string]interface{}) ([]byte, error) {
	// pongo2 を利用して、テンプレートファイルとデータから HTML を生成している
	// 生成した HTML はバイトデータに変えてreturn
	return pongo2.Must(pongo2.FromCache(templPath + file)).ExecuteBytes(data)
}

func render(c echo.Context, file string, data map[string]interface{}) error {
	b, err := htmlBlob(file,data)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.HTMLBlob(http.StatusOK, b)
}
