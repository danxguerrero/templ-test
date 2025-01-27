package handlers

import (
	"net/http"
	"strconv"
	"github.com/danxguerrero/templ-test/data"
	"github.com/danxguerrero/templ-test/templates"

	"github.com/labstack/echo/v4"
)

func ShowIndex(c echo.Context) error {
	return templates.Index(data.Tasks).Render(c.Request().Context(), c.Response().Writer)
}

func AddTask(c echo.Context) error {
	title := c.FormValue("title")
	data.AddTask(title)
	newTask := data.Tasks[len(data.Tasks)-1]
	return templates.Task(newTask.ID, newTask.Title, newTask.Completed).Render(c.Request().Context(), c.Response().Writer)
}

func ToggleTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	data.ToggleTask(id)
	for _, task := range data.Tasks {
		if task.ID == id {
			return templates.Task(task.ID, task.Title, task.Completed).Render(c.Request().Context(), c.Response().Writer)
		}
	}
	return c.NoContent(http.StatusNotFound)
}

func DeleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	data.DeleteTask(id)
	return c.NoContent(http.StatusOK)
}