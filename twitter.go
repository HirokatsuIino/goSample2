package main

import (
	"github.com/labstack/echo"
	"net/http"
	_ "net/http"
)
import _ "github.com/carlescere/scheduler"
import _ "github.com/labstack/echo"

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
