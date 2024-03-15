package main

import (
	"github.com/labstack/echo/v4"
	"github.com/you-5805/learning-go/internal/delivery"
	"github.com/you-5805/learning-go/internal/usecase"
)

func main() {
	e := echo.New()

	postUsecase := usecase.NewPostUsecase()
	postHandler := delivery.NewPostHandler(postUsecase)

	e.GET("/posts", func(c echo.Context) error {
		postHandler.GetPosts(c)
		return nil
	})

	e.Logger.Fatal(e.Start(":8080"))
}
