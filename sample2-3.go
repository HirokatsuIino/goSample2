package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type Message struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type Response struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:message`
	Stusts  string `json:status`
}

func main() {
	e := echo.New()
	e.POST("/send", sendMessage)
	e.Logger.Fatal(e.Start(":1323"))
}

func sendMessage(c echo.Context) error {
	m := new(Message)
	if error := c.Bind(m); error != nil {
		return error
	}
	r := new(Response)
	r.Name = m.Name
	r.Email = m.Email
	r.Message = m.Message
	r.Stusts = "success"
	return c.JSON(http.StatusOK, r)
}
