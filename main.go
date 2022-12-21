package main

import (
	"net/http"
	"strconv"
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
	e.GET("/new", articleNew)
	// :idはパラメーターで　任意の文字列を受け取る
  e.GET("/:id", articleShow)
  e.GET("/:id/edit", articleEdit)

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

	// `src/css` ディレクトリ配下のファイルに `/css` のパスでアクセスできるようにする
	e.Static("/css", "src/css")

  // アプリケーションインスタンスを返却
  return e
}

func articleIndex(c echo.Context) error {
	// ステータスコード 200 で、"Hello, World!" という文字列をレスポンス
	// return c.String(http.StatusOK, "Hello, World!")

	data := map[string]interface{}{
		"Message": "Article Index",
		"Now": time.Now(),
	}
	return render(c, "article/index.html", data)
}

func articleNew(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Article New",
		"Now": time.Now(),
	}
	return render(c, "article/new.html", data)
}

func articleShow(c echo.Context) error {
	// c.Param("id") でパスパラメータの :id の部分を抽出しています。
	// 例えば http://localhost:8080/999 でアクセスがあった場合、c.Param("id") によって 999 が取り出せます。
	// c.Param() で取り出した値は文字列型になるので、strconv パッケージの Atoi() 関数を使って数値型にキャストしています。
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Message": "Article Show",
		"Now": time.Now(),
		"ID": id,
	}
	return render(c, "article/show.html", data)
}


func articleEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data := map[string]interface{}{
		"Message": "Article Edit",
		"Now": time.Now(),
		"ID": id,
	}
	return render(c, "article/edit.html", data)
	
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
