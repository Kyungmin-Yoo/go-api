package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID string `param:"id" query:"id" header:"id" form:"id" json:"id" xml:"id"`
}

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World\n")
	})
	e.GET("/users/:id", getUser)
	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	return c.JSON(http.StatusOK, new(User))
}
