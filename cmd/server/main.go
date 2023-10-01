package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/posts/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, "Post ID: "+id)
	})

	e.Start(":8080")
}
