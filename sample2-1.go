package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {

	e := echo.New()
	e.GET("/users/:name", getUsername)
	e.Logger.Fatal(e.Start(":1323"))
}

func getUsername(c echo.Context) error {
	name := c.Param("name")

	return c.String(http.StatusOK, name)

}

