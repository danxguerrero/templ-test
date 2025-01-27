package main

import (
	"net/http"


	"github.com/danxguerrero/templ-test/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/static", "static")
	e.GET("/", handlers.ShowIndex)
	e.POST("/add", handlers.AddTask)
	e.POST("/toggle", handlers.ToggleTask)
	e.POST("/delete", handlers.DeleteTask)

	e.Logger.Fatal(e.Start(":8080"))
}